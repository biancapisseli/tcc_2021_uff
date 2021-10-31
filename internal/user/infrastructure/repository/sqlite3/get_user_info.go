package userreposqlite3

import (
	"context"
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/sqlite3"
	"ifoodish-store/pkg/sqlxtx"
	"net/http"

	"github.com/carlmjohnson/resperr"
)

func (r UserSQLite3Repository) GetUserInfo(
	ctx context.Context,
	userID uservo.UserID,
	addressID uservo.AddressID,
) (userInfo userent.RegisteredUser, err error) {
	tx, err := sqlxtx.GetTransaction(ctx)
	if err != nil {
		return userInfo, fmt.Errorf(
			"trying to get transaction to get sqlite3 user info: %w",
			err,
		)
	}

	if err := tx.Get(&userInfo, "SELECT id, email, name, phone FROM user WHERE id=$1", userID); err != nil {
		if sqlite3.IsErrNoRows(err) {
			return userInfo, resperr.WithCodeAndMessage(
				fmt.Errorf("error trying to get user info from sqlite3: %w", err),
				http.StatusNotFound,
				"Usuário não encontrado",
			)
		}
		return userInfo, resperr.WithStatusCode(
			fmt.Errorf("error trying to get user info from sqlite3: %w", err),
			http.StatusInternalServerError,
		)
	}

	return userInfo, nil
}
