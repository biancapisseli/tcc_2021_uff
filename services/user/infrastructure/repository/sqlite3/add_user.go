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
	"github.com/mattn/go-sqlite3"
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

	toAdd := struct {
		userent.User
		UserID   uservo.UserID          `json:"user_id"`
		Password uservo.PasswordEncoded `json:"password"`
	}{
		User:     user,
		UserID:   userID,
		Password: password,
	}
	if _, err := tx.NamedExec(
		`INSERT INTO user(
			id, email, name, password, phone
		) VALUES (
			:user_id, :email, :name, :password, :phone
		)`,
		toAdd,
	); err != nil {
		if errors.Is(err, sqlite3.ErrConstraintUnique) {
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
