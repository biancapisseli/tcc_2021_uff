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

func (r UserSQLite3Repository) GetUserAddress(
	ctx context.Context,
	userID uservo.UserID,
	addressID uservo.AddressID,
) (address userent.RegisteredAddress, err error) {
	tx, err := sqlxtx.GetTransaction(ctx)
	if err != nil {
		return address, fmt.Errorf(
			"trying to get transaction to get sqlite3 user address: %w",
			err,
		)
	}

	if err := tx.Get(&address, `
	SELECT
		id, street, district, city, state, complement,
		number, zipcode, latitude, longitude
	FROM address
	WHERE user_id=$1 AND id=$2
	`,
		userID,
		addressID,
	); err != nil {
		if sqlite3.IsErrNoRows(err) {
			return address, resperr.WithCodeAndMessage(
				fmt.Errorf("trying to get user info from sqlite3: %w", err),
				http.StatusNotFound,
				"Endereço não encontrado",
			)
		}
		return address, resperr.WithStatusCode(
			fmt.Errorf("trying to get userinfo from sqlite3: %w", err),
			http.StatusInternalServerError,
		)
	}

	return address, nil
}
