package userreposqlite3

import (
	"context"
	"fmt"
	"ifoodish-store/pkg/sqlite3"
	"ifoodish-store/pkg/sqlxtx"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
	"net/http"

	"github.com/carlmjohnson/resperr"
)

func (r UserSQLite3Repository) AddUserAddress(
	ctx context.Context,
	userID uservo.UserID,
	address userent.Address,
) (addressID uservo.AddressID, err error) {

	tx, err := sqlxtx.GetTransaction(ctx)
	if err != nil {
		return addressID, fmt.Errorf(
			"trying to get transaction to add new user address to sqlite3 db: %w",
			err,
		)
	}

	addressID = uservo.GenerateNewAddressID()

	toAdd := struct {
		userent.Address
		AddressID uservo.AddressID `json:"address_id"`
		UserID    uservo.UserID    `json:"user_id"`
	}{address, addressID, userID}

	if _, err = tx.NamedExec(
		`INSERT INTO address(
			id, street, district, city, state,
			complement, number, zipcode, user_id,
			latitude, longitude
		) VALUES (
			:address_id, :street, :district, :city, :state,
			:complement, :number, :zipcode, :user_id,
			:latitude, :longitude
		)`,
		toAdd,
	); err != nil {
		if sqlite3.IsForeignKeyErr(err) {
			return addressID, resperr.WithCodeAndMessage(
				fmt.Errorf("trying to add new user address to sqlite3 db: %w", err),
				http.StatusNotFound,
				"Usuário não encontrado",
			)
		}
		return addressID, resperr.WithStatusCode(
			fmt.Errorf("trying to add new user address to sqlite3 db: %w", err),
			http.StatusInternalServerError,
		)
	}

	return addressID, nil
}
