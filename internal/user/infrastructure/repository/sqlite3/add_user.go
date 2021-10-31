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

func (r UserSQLite3Repository) AddUser(
	ctx context.Context,
	user userent.User,
	password uservo.PasswordEncoded,
) (userID uservo.UserID, err error) {

	tx, err := sqlxtx.GetTransaction(ctx)
	if err != nil {
		return userID, fmt.Errorf(
			"trying to get transaction to add new user to sqlite3 db: %w",
			err,
		)
	}

	userID = uservo.GenerateNewUserID()

	newUser := userent.RegisteredUser{
		User: user,
		ID:   userID,
	}
	if _, err := tx.NamedExec(
		`INSERT INTO user(
			id, email, name, password, phone
		) VALUES (
			:id, :email, :name, :password, :phone
		)`,
		newUser,
	); err != nil {
		if sqlite3.IsUniqueErr(err) {
			return userID, resperr.WithCodeAndMessage(
				fmt.Errorf("adding new user to sqlite3 db: %w", err),
				http.StatusConflict,
				"O email já está em uso",
			)
		}
		return userID, resperr.WithStatusCode(
			fmt.Errorf("adding new user to sqlite3 db: %w", err),
			http.StatusInternalServerError,
		)
	}
	return userID, nil
}
