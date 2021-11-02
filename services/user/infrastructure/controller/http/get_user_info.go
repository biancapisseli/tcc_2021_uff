package userhttpcontroller

import (
	"fmt"
	userent "ifoodish-store/services/user/domain/entity"
)

func (c UserHTTPController) GetUserInfo(req Request) (
	user userent.RegisteredUser,
	err error,
) {
	userID, err := req.GetUserID()
	if err != nil {
		return user, fmt.Errorf("failed to get user id: %w", err)
	}

	user, err = c.useCases.GetUserInfo(
		req.Context(),
		userID,
	)
	if err != nil {
		return user, fmt.Errorf("failed use case: %w", err)
	}

	return user, nil
}
