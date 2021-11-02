package userreposqlite3

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"ifoodish-store/pkg/sqlxtx"
	uservo "ifoodish-store/services/user/domain/valueobject"

	"github.com/carlmjohnson/resperr"
)

func (r UserSQLite3Repository) UpdatePassword(
	ctx context.Context,
	userID uservo.UserID,
	newPassword uservo.PasswordEncoded,
) (err error) {
	tx, err := sqlxtx.GetTransaction(ctx)
	if err != nil {
		return fmt.Errorf(
			"trying to get transaction to update sqlite3 user password: %w",
			err,
		)
	}

	result, err := tx.Exec(
		"UPDATE user SET password=$1 WHERE id=$2",
		newPassword.String(),
		userID,
	)
	if err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf("trying to update user password from sqlite3: %w", err),
			http.StatusInternalServerError,
		)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf("trying to get rows affected from update user password from sqlite3: %w", err),
			http.StatusInternalServerError,
		)
	}

	if affectedRows == 0 {
		return resperr.WithCodeAndMessage(
			errors.New("affectedRows=0 on updating user password from sqlite3"),
			http.StatusNotFound,
			"Usuário não encontrado",
		)
	}

	return nil
}
