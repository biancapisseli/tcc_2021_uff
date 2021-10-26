package userhttpginctl

import (
	userctlint "ifoodish-store/internal/user/infrastructure/controller/interfaces"

	"github.com/labstack/echo/v4"
)

type UserHTTPGinController struct {
	useCases userctlint.UserUseCases
}

func New(useCases userctlint.UserUseCases) *UserHTTPGinController {
	return &UserHTTPGinController{
		useCases: useCases,
	}
}

func (c UserHTTPGinController) Register(router *echo.Group) {

	router.POST("/user", c.RegisterUser)
	router.GET("/user/:user_id", c.GetUserInfo)
	router.PUT("/user/:user_id", c.UpdateUserInfo)
	router.PUT("/user/:user_id/update_password", c.ChangePassword)

	router.POST("/user/:user_id/address", c.AddUserAddress)
	router.GET("/user/:user_id/address/:address_id", c.GetUserAddress)
	router.GET("/user/:user_id/address", c.GetUserAddresses)
	router.PUT("/user/:user_id/address/:address_id", c.UpdateUserAddress)
	router.DELETE("/user/:user_id/address/:address_id", c.RemoveUserAddress)
}
