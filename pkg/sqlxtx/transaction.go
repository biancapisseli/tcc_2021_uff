package sqlxtx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/carlmjohnson/resperr"
)

type Transaction interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type TransactionFinisher interface {
	Transaction
	Commit() (err error)
	Rollback() (err error)
}

type Transactioner interface {
	BeginTransaction() (tx TransactionFinisher, err error)
}

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
			errors.New("transaction doesn't exists"),
			http.StatusInternalServerError,
		)
	}

	tx, ok := interfaceValue.(TransactionFinisher)
	if !ok {
		return nil, resperr.WithStatusCode(
			errors.New("transaction with incorrect type"),
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
