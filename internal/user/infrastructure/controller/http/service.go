package userhttpctl

import (
	userrepo "ifoodish-store/internal/domain/repository"
)

type UserHTTPController struct {
	repo            userrepo.UserRepository
	passwordEncoder userrepo.PasswordEncoder
}

func New(repo userrepo.UserRepository, passwordEncoder userrepo.PasswordEncoder) *UserHTTPController {
	return &UserHTTPController{
		repo:            repo,
		passwordEncoder: passwordEncoder,
	}
}
