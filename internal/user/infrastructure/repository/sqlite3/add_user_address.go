package userreposqlite3

import (
	"context"
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/resperr"
	"ifoodish-store/pkg/sqlite3"
	"ifoodish-store/pkg/sqlxtx"
	"net/http"
)

func (r UserSQLite3Repository) AddUserAddress(
	ctx context.Context,
	userID uservo.UserID,
	address userent.Address,
) (addresID uservo.AddressID, err error) {

	tx, err := sqlxtx.GetTransaction(ctx)
	if err != nil {
		return 0, fmt.Errorf(
			"trying to get transaction to add new user address to sqlite3 db: %w",
			err,
		)
	}

	toAdd := struct {
		userent.Address
		UserID uservo.UserID `json:"user_id"`
	}{address, userID}

	result, err := tx.NamedExec(
		`INSERT INTO address(
			street, district, city, state,
			complement, number, zipcode, user_id,
			latitude, longitude
		) VALUES (
			:street, :district, :city, :state,
			:complement, :number, :zipcode, :user_id,
			:latitude, :longitude
		)`,
		toAdd,
	)
	if err != nil {
		if sqlite3.IsForeignKeyErr(err) {
			return 0, resperr.WithCodeAndMessage(
				fmt.Errorf("trying to add new user address to sqlite3 db: %w", err),
				http.StatusNotFound,
				"Usuário não encontrado",
			)
		}
		return 0, resperr.WithStatusCode(
			fmt.Errorf("trying to add new user address to sqlite3 db: %w", err),
			http.StatusInternalServerError,
		)
	}
	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, resperr.WithStatusCode(
			fmt.Errorf("trying to get new user address inserted ID: %w", err),
			http.StatusInternalServerError,
		)
	}

	return uservo.AddressID(insertedID), nil
}
