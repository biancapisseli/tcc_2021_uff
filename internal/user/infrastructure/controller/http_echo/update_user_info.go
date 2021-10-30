package userhttpechoctl

import (
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	"ifoodish-store/pkg/jwt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c UserHTTPGinController) UpdateUserInfo(echoCtx echo.Context) (err error) {

	reqCtx := echoCtx.Request().Context()

	userID, err := jwt.GetUserID(reqCtx)
	if err != nil {
		return fmt.Errorf("failed to get user id: %w", err)
	}

	var body userent.User
	if err := echoCtx.Bind(&body); err != nil {
		return fmt.Errorf("failed binding request body: %w", err)
	}

	user, err := userent.NewRegisteredUser(userent.RegisteredUser{
		ID:   userID,
		User: userent.User(body),
	})
	if err != nil {
		return fmt.Errorf("invalid user: %w", err)
	}

	if err := c.useCases.UpdateUserInfo(reqCtx, user); err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return echoCtx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Usuário atualizado com sucesso",
	})
}
