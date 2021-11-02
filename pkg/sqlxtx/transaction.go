package sqlxtx

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/carlmjohnson/resperr"
)

var (
	ErrTransactionNotFound      = errors.New("transaction not found on context")
	ErrTransactionIncorrectType = errors.New("transaction found with incorrect type")
)

type transactionKey struct{}

func BeginTransaction(db Transactioner, ctx context.Context) (context.Context, error) {
	tx, err := db.BeginTransaction()
	if err != nil {
		return ctx, resperr.WithStatusCode(
			fmt.Errorf("error trying to begin a transaction: %w", err),
			http.StatusInternalServerError,
		)
	}

	return context.WithValue(ctx, transactionKey{}, tx), nil
}

func getTransactionFinisher(ctx context.Context) (TransactionFinisher, error) {
	interfaceValue := ctx.Value(transactionKey{})
	if interfaceValue == nil {
		return nil, resperr.WithStatusCode(
			ErrTransactionNotFound,
			http.StatusInternalServerError,
		)
	}

	tx, ok := interfaceValue.(TransactionFinisher)
	if !ok {
		return nil, resperr.WithStatusCode(
			ErrTransactionIncorrectType,
			http.StatusInternalServerError,
		)
	}

	return tx, nil
}

func GetTransaction(ctx context.Context) (Transaction, error) {
	return getTransactionFinisher(ctx)
}

func CommitTransaction(ctx context.Context) error {
	tx, err := getTransactionFinisher(ctx)
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
	tx, err := getTransactionFinisher(ctx)
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
