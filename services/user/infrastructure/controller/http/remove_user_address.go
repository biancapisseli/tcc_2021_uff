package userhttpcontroller

import (
	"fmt"
	uservo "ifoodish-store/services/user/domain/valueobject"
	"net/http"

	"github.com/carlmjohnson/resperr"
)

func (c UserHTTPController) RemoveUserAddress(req Request) (err error) {

	userID, err := req.GetUserID()
	if err != nil {
		return fmt.Errorf("failed to get user id: %w", err)
	}

	var uri struct {
		AddressID int64 `url:"address_id"`
	}
	if err := req.ParseURLParams(&uri); err != nil {
		return resperr.WithCodeAndMessage(
			fmt.Errorf("failed binding request uri: %w", err),
			http.StatusBadRequest,
			"os parametros da URL estão incorretos",
		)
	}

	addressID, err := uservo.NewAddressID(uri.AddressID)
	if err != nil {
		return fmt.Errorf("invalid address id: %w", err)
	}

	if err := c.useCases.RemoveUserAddress(req.Context(), userID, addressID); err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return nil
}
