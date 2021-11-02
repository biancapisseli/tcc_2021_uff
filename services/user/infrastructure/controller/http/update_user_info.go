package userhttpcontroller

import (
	"fmt"
	userent "ifoodish-store/services/user/domain/entity"
)

func (c UserHTTPController) UpdateUserInfo(req Request) (err error) {

	userID, err := req.GetUserID()
	if err != nil {
		return fmt.Errorf("failed to get user id: %w", err)
	}

	var body userent.User
	if err := req.ParseBody(&body); err != nil {
		return fmt.Errorf("failed binding request body: %w", err)
	}

	user, err := userent.NewRegisteredUser(userID.String(), body)
	if err != nil {
		return fmt.Errorf("invalid user: %w", err)
	}

	if err := c.useCases.UpdateUserInfo(req.Context(), user); err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return nil
}
