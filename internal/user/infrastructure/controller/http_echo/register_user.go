package userhttpechoctl

import (
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c UserHTTPGinController) RegisterUser(echoCtx echo.Context) (err error) {

	type passwordRawClone uservo.PasswordRaw
	type userClone userent.User
	var body struct {
		userClone
		Password        passwordRawClone `json:"password"`
		PasswordConfirm passwordRawClone `json:"password_confirm"`
	}

	if err := echoCtx.Bind(&body); err != nil {
		return fmt.Errorf("failed binding request body: %w", err)
	}

	password, err := uservo.NewPasswordRaw(
		uservo.PasswordRaw(body.Password).String(),
	)
	if err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}

	passwordConfirm, err := uservo.NewPasswordRaw(
		uservo.PasswordRaw(body.PasswordConfirm).String(),
	)
	if err != nil {
		return fmt.Errorf("invalid password confirm: %w", err)
	}

	user, err := userent.NewUser(userent.User(body.userClone))
	if err != nil {
		return fmt.Errorf("invalid user: %w", err)
	}

	userID, err := c.useCases.RegisterUser(
		echoCtx.Request().Context(),
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
