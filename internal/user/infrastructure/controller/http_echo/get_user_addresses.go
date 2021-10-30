package userhttpechoctl

import (
	"fmt"
	"ifoodish-store/pkg/jwt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c UserHTTPGinController) GetUserAddresses(echoCtx echo.Context) (err error) {

	reqCtx := echoCtx.Request().Context()

	userID, err := jwt.GetUserID(reqCtx)
	if err != nil {
		return fmt.Errorf("failed to get user id: %w", err)
	}

	addresses, err := c.useCases.GetUserAddresses(reqCtx, userID)
	if err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return echoCtx.JSON(http.StatusOK, addresses)
}
