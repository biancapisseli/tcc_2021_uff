package usersvc

import (
	userent "ifoodish-store/internal/domain/entity"
)

type UserService struct {
	repo            userent.UserRepository
	passwordEncoder userent.PasswordEncoder
}

func New(repo userent.UserRepository, passwordEncoder userent.PasswordEncoder) *UserService {
	return &UserService{
		repo:            repo,
		passwordEncoder: passwordEncoder,
	}
}
