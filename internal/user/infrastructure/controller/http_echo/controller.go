package userhttpechoctl

import (
	userctlint "ifoodish-store/internal/user/infrastructure/controller/interfaces"
	"ifoodish-store/pkg/middleware"

	"github.com/labstack/echo/v4"
)

type UserHTTPGinController struct {
	useCases              userctlint.UserUseCases
	transactionMiddleware middleware.TransactionMiddleware
	jwtMiddleware         middleware.JWTHeaderMiddleware
}

func New(
	useCases userctlint.UserUseCases,
	transactionMiddleware middleware.TransactionMiddleware,
	jwtMiddleware middleware.JWTHeaderMiddleware,
) *UserHTTPGinController {
	return &UserHTTPGinController{
		useCases:              useCases,
		transactionMiddleware: transactionMiddleware,
		jwtMiddleware:         jwtMiddleware,
	}
}

func (c UserHTTPGinController) Register(router *echo.Group) {

	router.Use(c.transactionMiddleware.Middleware)
	router.POST("/register", c.RegisterUser)

	authRouter := router.Group("", c.jwtMiddleware.Middleware)

	authRouter.GET("/", c.GetUserInfo)
	authRouter.PUT("/", c.UpdateUserInfo)
	authRouter.PUT("/update_password", c.ChangePassword)

	authRouter.POST("/address", c.AddUserAddress)
	authRouter.GET("/address/:address_id", c.GetUserAddress)
	authRouter.GET("/address", c.GetUserAddresses)
	authRouter.PUT("/address", c.UpdateUserAddress)
	authRouter.DELETE("/address/:address_id", c.RemoveUserAddress)
}
