package userhttpechoctl

import (
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/resperr"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c UserHTTPGinController) UpdateUserAddress(echoCtx echo.Context) (err error) {
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

	type addressClone userent.Address
	var body addressClone

	if err := echoCtx.Bind(&body); err != nil {
		return fmt.Errorf("failed binding request body: %w", err)
	}

	address, err := userent.NewRegisteredAddress(userent.RegisteredAddress{
		ID:      addressID,
		Address: userent.Address(body),
	})
	if err != nil {
		return fmt.Errorf("invalid address: %w", err)
	}

	if err := c.useCases.UpdateUserAddress(
		echoCtx.Request().Context(),
		userID,
		address,
	); err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return echoCtx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Endereço atualizado com sucesso",
	})
}
