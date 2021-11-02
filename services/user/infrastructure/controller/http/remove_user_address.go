package userhttpcontroller

import (
	"fmt"
	uservo "ifoodish-store/services/user/domain/valueobject"
)

func (c UserHTTPController) RemoveUserAddress(req Request) (err error) {

	userID, err := req.GetUserID()
	if err != nil {
		return fmt.Errorf("failed to get user id: %w", err)
	}

	addressID, err := uservo.NewAddressID(req.GetURLParam("address_id"))
	if err != nil {
		return fmt.Errorf("invalid address id: %w", err)
	}

	if err := c.useCases.RemoveUserAddress(req.Context(), userID, addressID); err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return nil
}
