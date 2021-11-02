package userhttpcontroller

import (
	"encoding/json"
	"fmt"
	uservo "ifoodish-store/services/user/domain/valueobject"
)

type ChangePasswordBody struct {
	CurrentPassword    uservo.PasswordRaw `json:"current_password"`
	NewPassword        uservo.PasswordRaw `json:"new_password"`
	NewPasswordConfirm uservo.PasswordRaw `json:"new_password_confirm"`
}

func (b *ChangePasswordBody) UnmarshalJSON(data []byte) error {
	type clone ChangePasswordBody
	var bodyClone clone

	if err := json.Unmarshal(data, &bodyClone); err != nil {
		return fmt.Errorf("failed to bind request body: %w", err)
	}

	currentPassword, err := uservo.NewPasswordRaw(bodyClone.CurrentPassword.String())
	if err != nil {
		return fmt.Errorf("failed to bind request body: %w", err)
	}

	newPassword, err := uservo.NewPasswordRaw(bodyClone.NewPassword.String())
	if err != nil {
		return fmt.Errorf("failed to bind request body: %w", err)
	}

	newPasswordConfirm, err := uservo.NewPasswordRaw(bodyClone.NewPasswordConfirm.String())
	if err != nil {
		return fmt.Errorf("failed to bind request body: %w", err)
	}

	*b = ChangePasswordBody{
		CurrentPassword:    currentPassword,
		NewPassword:        newPassword,
		NewPasswordConfirm: newPasswordConfirm,
	}

	return nil
}

func (c UserHTTPController) ChangePassword(req Request) (err error) {

	userID, err := req.GetUserID()
	if err != nil {
		return fmt.Errorf("failed to get user id: %w", err)
	}

	var body ChangePasswordBody
	if err := req.ParseBodyParams(&body); err != nil {
		return fmt.Errorf("failed to bind request body: %w", err)
	}

	if err := c.useCases.ChangePassword(
		req.Context(),
		userID,
		body.CurrentPassword,
		body.NewPassword,
		body.NewPasswordConfirm,
	); err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return nil
}
