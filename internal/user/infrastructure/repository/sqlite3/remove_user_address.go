package userreposqlite3

import (
	"context"
	"errors"
	"fmt"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/sqlxtx"
	"net/http"

	"github.com/carlmjohnson/resperr"
)

func (r UserSQLite3Repository) RemoveUserAddress(
	ctx context.Context,
	userID uservo.UserID,
	addressID uservo.AddressID,
) (err error) {
	tx, err := sqlxtx.GetTransaction(ctx)
	if err != nil {
		return fmt.Errorf(
			"trying to get transaction to delete sqlite3 user address: %w",
			err,
		)
	}
	result, err := tx.Exec(
		"DELETE FROM address WHERE id=$1 AND user_id=$2",
		addressID,
		userID,
	)
	if err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf("trying to delete user address from sqlite3: %w", err),
			http.StatusInternalServerError,
		)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf("trying to get rows affected from deleted user address from sqlite3: %w", err),
			http.StatusInternalServerError,
		)
	}
	if affectedRows == 0 {
		return resperr.WithCodeAndMessage(
			errors.New("affectedRows=0 on removing user address from sqlite3"),
			http.StatusNotFound,
			"Endereço não encontrado",
		)
	}

	return nil
}
