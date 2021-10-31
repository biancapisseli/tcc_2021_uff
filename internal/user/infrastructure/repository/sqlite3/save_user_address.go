package userreposqlite3

import (
	"context"
	"errors"
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/sqlxtx"
	"net/http"

	"github.com/carlmjohnson/resperr"
)

func (r UserSQLite3Repository) SaveUserAddress(
	ctx context.Context,
	userID uservo.UserID,
	address userent.RegisteredAddress,
) (err error) {
	tx, err := sqlxtx.GetTransaction(ctx)
	if err != nil {
		return fmt.Errorf(
			"trying to get transaction to save sqlite3 user address: %w",
			err,
		)
	}

	toUpdate := struct {
		userent.RegisteredAddress
		UserID uservo.UserID `json:"user_id"`
	}{address, userID}
	result, err := tx.NamedExec(`
	UPDATE address SET
		street=:street,
		district=:district,
		city=:city,
		state=:state,
		complement=:complement,
		number=:number,
		zipcode=:zipcode,
		latitude=:latitude,
		longitude=:longitude
	WHERE id=:id AND user_id=:user_id
	`, toUpdate)

	if err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf("trying to save user address to sqlite3: %w", err),
			http.StatusInternalServerError,
		)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf(
				"trying to get rows affected while saving user address to sqlite3: %w",
				err,
			),
			http.StatusInternalServerError,
		)
	}

	if affectedRows == 0 {
		return resperr.WithCodeAndMessage(
			errors.New("affectedRows=0 on saving user address to sqlite3"),
			http.StatusNotFound,
			"Endereço não encontrado",
		)
	}

	return nil
}
