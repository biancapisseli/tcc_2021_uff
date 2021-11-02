package userhttpcontroller

import (
	"fmt"
	userent "ifoodish-store/services/user/domain/entity"
)

func (c UserHTTPController) GetUserAddresses(req Request) (
	addresses []userent.RegisteredAddress,
	err error,
) {

	userID, err := req.GetUserID()
	if err != nil {
		return nil, fmt.Errorf("failed to get user id: %w", err)
	}

	addresses, err = c.useCases.GetUserAddresses(req.Context(), userID)
	if err != nil {
		return nil, fmt.Errorf("failed use case: %w", err)
	}

	return addresses, nil
}
