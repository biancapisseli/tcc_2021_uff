package userhttpcontroller

import (
	"fmt"
	userent "ifoodish-store/services/user/domain/entity"
)

func (c UserHTTPController) UpdateUserAddress(req Request) (err error) {

	userID, err := req.GetUserID()
	if err != nil {
		return fmt.Errorf("failed to get user id: %w", err)
	}

	var body userent.RegisteredAddress
	if err := req.ParseBodyParams(&body); err != nil {
		return fmt.Errorf("failed binding request body: %w", err)
	}

	if err := c.useCases.UpdateUserAddress(req.Context(), userID, body); err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return nil
}
