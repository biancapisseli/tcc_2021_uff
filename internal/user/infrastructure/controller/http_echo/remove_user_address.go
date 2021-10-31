package userhttpechoctl

import (
	"fmt"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/jwt"
	"net/http"

	"github.com/carlmjohnson/resperr"

	"github.com/labstack/echo/v4"
)

func (c UserHTTPGinController) RemoveUserAddress(echoCtx echo.Context) (err error) {

	reqCtx := echoCtx.Request().Context()

	userID, err := jwt.GetUserID(reqCtx)
	if err != nil {
		return fmt.Errorf("failed to get user id: %w", err)
	}

	type addressIDClone uservo.AddressID
	var uri struct {
		AddressID addressIDClone `param:"address_id"`
	}
	if err := echoCtx.Bind(&uri); err != nil {
		return resperr.WithCodeAndMessage(
			fmt.Errorf("failed binding request uri: %w", err),
			http.StatusBadRequest,
			"os parametros da URL estão incorretos",
		)
	}

	addressID, err := uservo.NewAddressID(int64(uri.AddressID))
	if err != nil {
		return fmt.Errorf("invalid address id: %w", err)
	}

	if err := c.useCases.RemoveUserAddress(reqCtx, userID, addressID); err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return echoCtx.JSON(http.StatusOK, map[string]interface{}{
		"message": "endereço removido com sucesso",
	})
}
