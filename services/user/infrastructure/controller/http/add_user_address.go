package userhttpcontroller

import (
	"fmt"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
)

func (c UserHTTPController) AddUserAddress(req Request) (
	addressID uservo.AddressID,
	err error,
) {

	userID, err := req.GetUserID()
	if err != nil {
		return addressID, fmt.Errorf("failed to get user id: %w", err)
	}

	var body userent.Address
	if err := req.ParseBodyParams(&body); err != nil {
		return addressID, fmt.Errorf("failed binding request body: %w", err)
	}

	addressID, err = c.useCases.AddUserAddress(req.Context(), userID, body)
	if err != nil {
		return addressID, fmt.Errorf("failed use case: %w", err)
	}

	return addressID, nil
}
