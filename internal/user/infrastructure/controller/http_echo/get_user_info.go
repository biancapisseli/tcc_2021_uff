package userhttpechoctl

import (
	"fmt"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/resperr"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c UserHTTPGinController) GetUserInfo(echoCtx echo.Context) (err error) {
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

	user, err := c.useCases.GetUserInfo(
		echoCtx.Request().Context(),
		userID,
	)
	if err != nil {
		return fmt.Errorf("failed use case: %w", err)
	}

	return echoCtx.JSON(http.StatusOK, user)
}
