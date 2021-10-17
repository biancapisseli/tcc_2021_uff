package usersvc

import (
	userrepo "ifoodish-store/internal/user/domain/repository"
)

type UserService struct {
	repo            userrepo.UserRepository
	passwordEncoder userrepo.PasswordEncoder
}

func New(repo userrepo.UserRepository, passwordEncoder userrepo.PasswordEncoder) *UserService {
	return &UserService{
		repo:            repo,
		passwordEncoder: passwordEncoder,
	}
}
