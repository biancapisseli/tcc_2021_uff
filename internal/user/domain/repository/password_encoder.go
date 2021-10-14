package userrepo

import (
	uservo "ifoodish-store/internal/domain/valueobject"
)

type PasswordEncoder interface {
	EncodePassword(rawPassword uservo.PasswordRaw) (encodedPassword uservo.PasswordEncoded, err error)
}
