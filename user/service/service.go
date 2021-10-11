package usersvc

import (
	userdom "ifoodish-store/user/domain"
)

type UserService struct {
	repo            userdom.UserRepository
	passwordEncoder userdom.PasswordEncoder
}

func New(repo userdom.UserRepository, passwordEncoder userdom.PasswordEncoder) *UserService {
	return &UserService{
		repo:            repo,
		passwordEncoder: passwordEncoder,
	}
}
