package userhttpcontroller

import (
	"encoding/json"
	"fmt"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
)

type RegisterUserBody struct {
	Password        uservo.PasswordRaw `json:"password"`
	PasswordConfirm uservo.PasswordRaw `json:"password_confirm"`
	User            userent.User       `json:"user"`
}

func (b *RegisterUserBody) UnmarshalJSON(data []byte) error {
	type clone RegisterUserBody
	var bodyClone clone

	if err := json.Unmarshal(data, &bodyClone); err != nil {
		return fmt.Errorf("failed to bind request body: %w", err)
	}

	password, err := uservo.NewPasswordRaw(bodyClone.Password.String())
	if err != nil {
		return fmt.Errorf("failed to bind request body: %w", err)
	}

	passwordConfirm, err := uservo.NewPasswordRaw(bodyClone.PasswordConfirm.String())
	if err != nil {
		return fmt.Errorf("failed to bind request body: %w", err)
	}

	*b = RegisterUserBody{
		User:            bodyClone.User,
		Password:        password,
		PasswordConfirm: passwordConfirm,
	}

	return nil
}

func (c UserHTTPController) RegisterUser(req Request) (
	userID uservo.UserID,
	err error,
) {
	var body RegisterUserBody
	if err := req.ParseBody(&body); err != nil {
		return userID, fmt.Errorf("failed binding request body: %w", err)
	}

	userID, err = c.useCases.RegisterUser(
		req.Context(),
		body.User,
		body.Password,
		body.PasswordConfirm,
	)
	if err != nil {
		return userID, fmt.Errorf("failed use case: %w", err)
	}

	return userID, nil
}
