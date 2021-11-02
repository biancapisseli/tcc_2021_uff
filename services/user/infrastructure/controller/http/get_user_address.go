package userhttpcontroller

import (
	"fmt"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
	"net/http"

	"github.com/carlmjohnson/resperr"
)

func (c UserHTTPController) GetUserAddress(req Request) (
	address userent.RegisteredAddress,
	err error,
) {

	userID, err := req.GetUserID()
	if err != nil {
		return address, fmt.Errorf("failed to get user id: %w", err)
	}

	type addressIDClone uservo.AddressID
	var uri struct {
		AddressID addressIDClone `url:"address_id"`
	}

	if err := req.ParseURLParams(&uri); err != nil {
		return address, resperr.WithCodeAndMessage(
			fmt.Errorf("failed binding request uri: %w", err),
			http.StatusBadRequest,
			"os parametros da URL estão incorretos",
		)
	}

	addressID, err := uservo.NewAddressID(int64(uri.AddressID))
	if err != nil {
		return address, fmt.Errorf("invalid address id: %w", err)
	}

	address, err = c.useCases.GetUserAddress(req.Context(), userID, addressID)
	if err != nil {
		return address, fmt.Errorf("failed use case: %w", err)
	}

	return address, nil
}
