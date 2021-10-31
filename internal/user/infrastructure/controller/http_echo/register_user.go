package userhttpechoctl

import (
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c UserHTTPGinController) RegisterUser(echoCtx echo.Context) (err error) {

	reqCtx := echoCtx.Request().Context()

	var user userent.User
	if err := echoCtx.Bind(&user); err != nil {
		return fmt.Errorf("failed binding request body: %w", err)
	}

	type passwordRawClone uservo.PasswordRaw
	var passwordParams struct {
		Password        passwordRawClone `json:"password"`
		PasswordConfirm passwordRawClone `json:"password_confirm"`
	}
	if err := echoCtx.Bind(&user); err != nil {
		return fmt.Errorf("failed binding request body: %w", err)
	}

	password, err := uservo.NewPasswordRaw(
		uservo.PasswordRaw(passwordParams.Password).String(),
	)
	if err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}

	passwordConfirm, err := uservo.NewPasswordRaw(
		uservo.PasswordRaw(passwordParams.PasswordConfirm).String(),
	)
	if err != nil {
		return fmt.Errorf("invalid password confirm: %w", err)
	}

	userID, err := c.useCases.RegisterUser(
		reqCtx,
		user,
		password,
		passwordConfirm,
	)
	if err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return echoCtx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Usuário registrado com sucesso",
		"user_id": userID,
	})
}
