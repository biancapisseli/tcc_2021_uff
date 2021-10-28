package userhttpechoctl

import (
	"fmt"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/resperr"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c UserHTTPGinController) ChangePassword(echoCtx echo.Context) (err error) {

	type userIDClone uservo.UserID
	var uri struct {
		UserID userIDClone `param:"user_id"`
	}
	if err := echoCtx.Bind(&uri); err != nil {
		return resperr.WithCodeAndMessage(
			fmt.Errorf("failed binding request uri: %w", err),
			http.StatusBadRequest,
			"os parametros da URL estão incorretos",
		)
	}

	userID, err := uservo.NewUserID(uservo.UserID(uri.UserID).String())
	if err != nil {
		return fmt.Errorf("invalid user id: %w", err)
	}

	type passwordRawClone uservo.PasswordRaw
	var body struct {
		CurrentPassword    passwordRawClone `json:"current_password"`
		NewPassword        passwordRawClone `json:"new_password"`
		NewPasswordConfirm passwordRawClone `json:"new_password_confirm"`
	}

	if err := echoCtx.Bind(&body); err != nil {
		return fmt.Errorf("failed to bind request body: %w", err)
	}

	currentPassword, err := uservo.NewPasswordRaw(
		uservo.PasswordRaw(body.CurrentPassword).String(),
	)
	if err != nil {
		return fmt.Errorf("invalid current password: %w", err)
	}

	newPassword, err := uservo.NewPasswordRaw(
		uservo.PasswordRaw(body.NewPassword).String(),
	)
	if err != nil {
		return fmt.Errorf("invalid new password: %w", err)
	}

	newPasswordConfirm, err := uservo.NewPasswordRaw(
		uservo.PasswordRaw(body.NewPasswordConfirm).String(),
	)
	if err != nil {
		return fmt.Errorf("invalid new password confirm: %w", err)
	}

	if err := c.useCases.ChangePassword(
		echoCtx.Request().Context(),
		userID,
		currentPassword,
		newPassword,
		newPasswordConfirm,
	); err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return echoCtx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Senha alterada com sucesso",
	})

}
