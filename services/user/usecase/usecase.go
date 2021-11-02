package useruc

import (
	userrepo "ifoodish-store/services/user/domain/repository"
	uservo "ifoodish-store/services/user/domain/valueobject"
)

type PasswordEncoder interface {
	EncodePassword(rawPassword uservo.PasswordRaw) (encodedPassword uservo.PasswordEncoded, err error)
}

type UserUseCases struct {
	repo            userrepo.UserRepository
	passwordEncoder PasswordEncoder
}

func New(repo userrepo.UserRepository, passwordEncoder PasswordEncoder) *UserUseCases {
	return &UserUseCases{
		repo:            repo,
		passwordEncoder: passwordEncoder,
	}
}
