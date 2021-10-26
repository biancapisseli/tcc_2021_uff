package middleware

import (
	"fmt"
	"ifoodish-store/pkg/sqlite3"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type TransactionMiddleware struct {
	conn *sqlite3.Connection
}

func NewTransactionMiddleware(conn *sqlite3.Connection) TransactionMiddleware {
	return TransactionMiddleware{
		conn: conn,
	}
}

func (m TransactionMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(echoCtx echo.Context) (err error) {

		parentCtx := echoCtx.Request().Context()

		txCtx, err := m.conn.BeginTransaction(parentCtx)
		if err != nil {
			return fmt.Errorf("error beginning transaction middleware: %w", err)
		}

		echoCtx.SetRequest(echoCtx.Request().WithContext(txCtx))

		defer func() {
			if err != nil {
				log.Warn(m.conn.RollbackTransaction(txCtx))
			}
		}()

		if err := next(echoCtx); err != nil {
			return err
		}

		return m.conn.CommitTransaction(txCtx)
	}
}
