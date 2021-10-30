package sqlite3

import (
	"database/sql"
	"fmt"
	"ifoodish-store/pkg/sqlitemigrator"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	sqlite3 "github.com/mattn/go-sqlite3"
)

const (
	DATABASE_BUSY_TIMEOUT = "5000"
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

func New(path string, migration map[string][]string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open(
		"sqlite3_with_fk",
		path+"?cache=shared&_busy_timeout="+DATABASE_BUSY_TIMEOUT,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating sqlite3 connection pool: %w", err)
	}

	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)

	db.SetMaxOpenConns(1)

	migrator := sqlitemigrator.New(db, migration)
	if err := migrator.Run(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
