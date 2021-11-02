package userhttpcontroller

import (
	usercontrollerinterface "ifoodish-store/services/user/infrastructure/controller/interface"
)

type UserHTTPController struct {
	useCases usercontrollerinterface.UserUseCases
}

func New(
	useCases usercontrollerinterface.UserUseCases,
) *UserHTTPController {
	return &UserHTTPController{
		useCases: useCases,
	}
}
