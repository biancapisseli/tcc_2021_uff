package userreposqlite3

import (
	"context"
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/sqlxtx"
	"net/http"

	"github.com/carlmjohnson/resperr"
)

func (r UserSQLite3Repository) GetUserAddresses(
	ctx context.Context,
	userID uservo.UserID,
) (adresses []userent.RegisteredAddress, err error) {
	tx, err := sqlxtx.GetTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"trying to get transaction to get sqlite3 user addresses: %w",
			err,
		)
	}

	addresses := []userent.RegisteredAddress{}

	if err := tx.Select(&addresses, `
	SELECT
		id, street, district, city, state, complement,
		number, zipcode, latitude, longitude
	FROM address
	WHERE user_id=$1
	`, userID); err != nil {
		return nil, resperr.WithStatusCode(
			fmt.Errorf("trying to get user addresses from sqlite3: %w", err),
			http.StatusInternalServerError,
		)
	}

	return addresses, nil
}
