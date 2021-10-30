package middleware

import (
	"fmt"
	"ifoodish-store/pkg/sqlxtx"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type TransactionMiddleware struct {
	conn *sqlx.DB
}

func NewTransactionMiddleware(conn *sqlx.DB) TransactionMiddleware {
	return TransactionMiddleware{
		conn: conn,
	}
}

func (m TransactionMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(echoCtx echo.Context) (err error) {

		parentCtx := echoCtx.Request().Context()

		txCtx, err := sqlxtx.BeginTransaction(m.conn, parentCtx)
		if err != nil {
			return fmt.Errorf("error beginning transaction middleware: %w", err)
		}

		echoCtx.SetRequest(echoCtx.Request().WithContext(txCtx))

		defer func() {
			if err != nil {
				log.Warn(sqlxtx.RollbackTransaction(txCtx))
			}
		}()

		if err := next(echoCtx); err != nil {
			return err
		}

		return sqlxtx.CommitTransaction(txCtx)
	}
}
