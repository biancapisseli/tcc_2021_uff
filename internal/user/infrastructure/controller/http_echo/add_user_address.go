package userhttpechoctl

import (
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/resperr"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c UserHTTPGinController) AddUserAddress(echoCtx echo.Context) (err error) {

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

	type addressClone userent.Address
	var body addressClone

	if err := echoCtx.Bind(&body); err != nil {
		return fmt.Errorf("failed binding request body: %w", err)
	}

	address, err := userent.NewAddress(userent.Address(body))
	if err != nil {
		return fmt.Errorf("invalid address: %w", err)
	}

	addressID, err := c.useCases.AddUserAddress(
		echoCtx.Request().Context(),
		userID,
		address,
	)
	if err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return echoCtx.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Endereço adicionado com sucesso",
		"address_id": addressID,
	})
}
