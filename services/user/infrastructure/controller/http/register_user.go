package userhttpcontroller

import (
	"fmt"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
)

func (c UserHTTPController) RegisterUser(req Request) (
	userID uservo.UserID,
	err error,
) {
	type passwordRawClone uservo.PasswordRaw
	var body struct {
		userent.User
		Password        passwordRawClone `json:"password"`
		PasswordConfirm passwordRawClone `json:"password_confirm"`
	}

	if err := req.ParseBodyParams(&body); err != nil {
		return userID, fmt.Errorf("failed binding request body: %w", err)
	}

	user := body.User

	password, err := uservo.NewPasswordRaw(
		uservo.PasswordRaw(body.Password).String(),
	)
	if err != nil {
		return userID, fmt.Errorf("invalid password: %w", err)
	}

	passwordConfirm, err := uservo.NewPasswordRaw(
		uservo.PasswordRaw(body.PasswordConfirm).String(),
	)
	if err != nil {
		return userID, fmt.Errorf("invalid password confirm: %w", err)
	}

	userID, err = c.useCases.RegisterUser(
		req.Context(),
		user,
		password,
		passwordConfirm,
	)
	if err != nil {
		return userID, fmt.Errorf("failed use case: %w", err)
	}

	return userID, nil
}
