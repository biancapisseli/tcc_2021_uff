package userhttpechoctl

import (
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/resperr"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c UserHTTPGinController) UpdateUserInfo(echoCtx echo.Context) (err error) {
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

	type userClone userent.User
	var body userClone

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

	if err := c.useCases.UpdateUserInfo(
		echoCtx.Request().Context(),
		user,
	); err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return echoCtx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Usuário atualizado com sucesso",
	})
}
