package userhttpechoctl

import (
	"fmt"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/resperr"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c UserHTTPGinController) RemoveUserAddress(echoCtx echo.Context) (err error) {
	type userIDClone uservo.UserID
	type addressIDClone uservo.AddressID
	var uri struct {
		UserID    userIDClone    `param:"user_id"`
		AddressID addressIDClone `param:"address_id"`
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

	addressID, err := uservo.NewAddressID(int64(uri.AddressID))
	if err != nil {
		return fmt.Errorf("invalid address id: %w", err)
	}

	if err := c.useCases.RemoveUserAddress(
		echoCtx.Request().Context(),
		userID,
		addressID,
	); err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return echoCtx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Endereço removido com sucesso",
	})
}
