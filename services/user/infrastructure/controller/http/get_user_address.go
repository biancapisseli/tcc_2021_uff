package userhttpcontroller

import (
	"fmt"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
)

func (c UserHTTPController) GetUserAddress(req Request) (
	address userent.RegisteredAddress,
	err error,
) {

	userID, err := req.GetUserID()
	if err != nil {
		return address, fmt.Errorf("failed to get user id: %w", err)
	}

	addressID, err := uservo.NewAddressID(req.GetURLParam("address_id"))
	if err != nil {
		return address, fmt.Errorf("invalid address id: %w", err)
	}

	address, err = c.useCases.GetUserAddress(req.Context(), userID, addressID)
	if err != nil {
		return address, fmt.Errorf("failed use case: %w", err)
	}

	return address, nil
}
