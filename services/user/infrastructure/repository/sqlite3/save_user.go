package userreposqlite3

import (
	"context"
	"errors"
	"fmt"
	"ifoodish-store/pkg/sqlxtx"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
	"net/http"

	"github.com/carlmjohnson/resperr"
)

func (r UserSQLite3Repository) SaveUser(
	ctx context.Context,
	userID uservo.UserID,
	user userent.User,
) (err error) {
	tx, err := sqlxtx.GetTransaction(ctx)
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
