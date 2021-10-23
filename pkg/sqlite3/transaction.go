package sqlite3

import (
	"context"
	"errors"
	"fmt"
	"ifoodish-store/pkg/resperr"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type transactionKey struct{}

func (c *Connection) BeginTransaction(ctx context.Context) (context.Context, error) {
	tx, err := c.db.Beginx()
	if err != nil {
		return ctx, resperr.WithStatusCode(
			fmt.Errorf("error trying to begin a transaction: %w", err),
			http.StatusInternalServerError,
		)
	}

	return context.WithValue(ctx, transactionKey{}, tx), nil
}

func (c *Connection) GetTransaction(ctx context.Context) (*sqlx.Tx, error) {
	interfaceValue := ctx.Value(transactionKey{})
	if interfaceValue == nil {
		return nil, resperr.WithStatusCode(
			errors.New("transaction doesn't exists"),
			http.StatusInternalServerError,
		)
	}

	tx, ok := interfaceValue.(*sqlx.Tx)
	if !ok {
		return nil, resperr.WithStatusCode(
			errors.New("transaction incorrect type"),
			http.StatusInternalServerError,
		)
	}

	return tx, nil
}

func (c *Connection) Commit(ctx context.Context) error {
	tx, err := c.GetTransaction(ctx)
	if err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf("error trying to commit transaction: %w", err),
			http.StatusInternalServerError,
		)
	}

	err = tx.Commit()
	if err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf("error trying to commit transaction: %w", err),
			http.StatusInternalServerError,
		)
	}

	return nil
}

func (c *Connection) Rollback(ctx context.Context) error {
	tx, err := c.GetTransaction(ctx)
	if err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf("error trying to rollback transaction: %w", err),
			http.StatusInternalServerError,
		)
	}

	err = tx.Rollback()
	if err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf("error trying to rollback transaction: %w", err),
			http.StatusInternalServerError,
		)
	}

	return nil
}
