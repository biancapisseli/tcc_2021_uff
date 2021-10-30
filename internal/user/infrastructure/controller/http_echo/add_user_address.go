package userhttpechoctl

import (
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	"ifoodish-store/pkg/jwt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c UserHTTPGinController) AddUserAddress(echoCtx echo.Context) (err error) {

	reqCtx := echoCtx.Request().Context()

	userID, err := jwt.GetUserID(reqCtx)
	if err != nil {
		return fmt.Errorf("failed to get user id: %w", err)
	}

	type addressClone userent.Address
	var body addressClone

	if err := echoCtx.Bind(&body); err != nil {
		return fmt.Errorf("failed binding request body: %w", err)
	}

	address, err := userent.NewAddress(userent.Address(body))
	if err != nil {
		return fmt.Errorf("invalid body: %w", err)
	}

	addressID, err := c.useCases.AddUserAddress(reqCtx, userID, address)
	if err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return echoCtx.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Endereço adicionado com sucesso",
		"address_id": addressID,
	})
}
