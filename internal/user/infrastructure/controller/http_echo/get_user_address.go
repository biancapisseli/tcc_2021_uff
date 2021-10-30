package userhttpechoctl

import (
	"fmt"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/jwt"
	"ifoodish-store/pkg/resperr"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c UserHTTPGinController) GetUserAddress(echoCtx echo.Context) (err error) {

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

	resp, err := c.useCases.GetUserAddress(reqCtx, userID, addressID)
	if err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return echoCtx.JSON(http.StatusOK, resp)
}
