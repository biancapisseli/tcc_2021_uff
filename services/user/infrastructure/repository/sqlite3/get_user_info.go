package userreposqlite3

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"ifoodish-store/pkg/sqlxtx"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
	"net/http"

	"github.com/carlmjohnson/resperr"
)

func (r UserSQLite3Repository) GetUserInfo(
	ctx context.Context,
	userID uservo.UserID,
) (userInfo userent.RegisteredUser, err error) {
	tx, err := sqlxtx.GetTransaction(ctx)
	if err != nil {
		return userInfo, fmt.Errorf(
			"trying to get transaction to get sqlite3 user info: %w",
			err,
		)
	}

	if err := tx.Get(&userInfo, "SELECT id, email, name, phone FROM user WHERE id=$1", userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
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
