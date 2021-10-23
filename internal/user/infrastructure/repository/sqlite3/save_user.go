package userreposqlite3

import (
	"context"
	"errors"
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	"ifoodish-store/pkg/resperr"
	"net/http"
)

func (r UserSQLite3Repository) SaveUser(
	ctx context.Context,
	user userent.RegisteredUser,
) (err error) {
	tx, err := r.db.GetTransaction(ctx)
	if err != nil {
		return fmt.Errorf(
			"trying to get transaction to save sqlite3 user: %w",
			err,
		)
	}

	results, err := tx.NamedExec(
		"UPDATE user SET name=:name, phone=:phone WHERE id=:id",
		user,
	)
	if err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf("trying to save user to sqlite3: %w", err),
			http.StatusInternalServerError,
		)
	}
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf("trying to get rows affected while saving user to sqlite3: %w", err),
			http.StatusInternalServerError,
		)
	}
	if rowsAffected == 0 {
		return resperr.WithCodeAndMessage(
			errors.New("rowsAffected=0 on saving user to sqlite3"),
			http.StatusNotFound,
			"Usuário não encontrado",
		)
	}

	return nil
}
