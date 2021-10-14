package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	sqlite3 "github.com/mattn/go-sqlite3"
)

func init() {
	sql.Register("sqlite3_with_fk",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				_, err := conn.Exec("PRAGMA foreign_keys = ON", nil)
				return err
			},
		})
}

type Connection struct {
	db sqlx.DB
}

type transactionKey struct{}

func New(path string, migration map[string][]string) (connection *Connection, err error) {
	connection = &Connection{}
	db, err := sqlx.Open(
		"sqlite3_with_fk",
		path+"?cache=shared&_busy_timeout="+DATABASE_BUSY_TIMEOUT,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating sqlite3 connection pool: %w", err)
	}

	db.SetMaxOpenConns(1)

	migrator := newMigrator(db, migration)
	if err := migrator.run(); err != nil {
		db.Close()
		return nil, err
	}

	return connection, nil
}

func (c *Connection) BeginTransaction(ctx context.Context) (context.Context, error) {
	tx, err := c.db.Beginx()
	if err != nil {
		return ctx, fmt.Errorf("error trying to begin a transaction: %w", err)
	}

	return context.WithValue(ctx, transactionKey{}, tx), nil
}

func (c *Connection) GetTransaction(ctx context.Context) (*sqlx.Tx, error) {
	interfaceValue := ctx.Value(transactionKey{})
	if interfaceValue == nil {
		return nil, errors.New("transaction doesn't exists")
	}

	tx, ok := interfaceValue.(*sqlx.Tx)
	if !ok {
		return nil, errors.New("transaction incorrect type")
	}

	return tx, nil
}

func (c *Connection) Commit(ctx context.Context) error {
	tx, err := c.GetTransaction(ctx)
	if err != nil {
		return fmt.Errorf("error trying to commit transaction: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error trying to commit transaction: %w", err)
	}

	return nil
}

func (c *Connection) Rollback(ctx context.Context) error {
	tx, err := c.GetTransaction(ctx)
	if err != nil {
		return fmt.Errorf("error trying to rollback transaction: %w", err)
	}

	err = tx.Rollback()
	if err != nil {
		return fmt.Errorf("error trying to rollback transaction: %w", err)
	}

	return nil
}
