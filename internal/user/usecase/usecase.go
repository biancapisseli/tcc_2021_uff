package useruc

import (
	userrepo "ifoodish-store/internal/user/domain/repository"
)

type UserUseCases struct {
	repo            userrepo.UserRepository
	passwordEncoder userrepo.PasswordEncoder
}

func New(repo userrepo.UserRepository, passwordEncoder userrepo.PasswordEncoder) *UserUseCases {
	return &UserUseCases{
		repo:            repo,
		passwordEncoder: passwordEncoder,
	}
}
