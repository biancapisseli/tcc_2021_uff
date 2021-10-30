package sqlxtx

import (
	"context"
	"errors"
	"fmt"
	"ifoodish-store/pkg/resperr"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type transactionKey struct{}

func BeginTransaction(db *sqlx.DB, ctx context.Context) (context.Context, error) {
	tx, err := db.Beginx()
	if err != nil {
		return ctx, resperr.WithStatusCode(
			fmt.Errorf("error trying to begin a transaction: %w", err),
			http.StatusInternalServerError,
		)
	}

	return context.WithValue(ctx, transactionKey{}, tx), nil
}

func GetTransaction(ctx context.Context) (*sqlx.Tx, error) {
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
			errors.New("transaction with incorrect type"),
			http.StatusInternalServerError,
		)
	}

	return tx, nil
}

func CommitTransaction(ctx context.Context) error {
	tx, err := GetTransaction(ctx)
	if err != nil {
		return fmt.Errorf("error trying to commit transaction: %w", err)

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

func RollbackTransaction(ctx context.Context) error {
	tx, err := GetTransaction(ctx)
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
