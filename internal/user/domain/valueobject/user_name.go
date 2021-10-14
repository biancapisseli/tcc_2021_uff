package uservo

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
)

const (
	minUserNameLength = 3
	maxUserNameLength = 50
)

var (
	ErrUserNameMinLength        = fmt.Errorf("nome de usuário deve possui mais de %d caracteres", minUserNameLength)
	ErrUserNameMaxLength        = fmt.Errorf("nome de usuário deve possui menos de %d caracteres", maxUserNameLength)
	ErrUserNameInvalidCharacter = fmt.Errorf("nome de usuário deve conter apenas letras e espaços")
)

type UserName string

func NewUserName(value string) (UserName, error) {
	if len(value) > maxUserNameLength {
		return "", ErrUserNameMaxLength
	}
	if len(value) < minUserNameLength {
		return "", ErrUserNameMinLength
	}

	if !valid.Matches(value, `^[a-zA-Z\s]*$`) {
		return "", ErrUserNameInvalidCharacter
	}
	return UserName(value), nil
}

func (u UserName) Equals(other UserName) bool {
	return u.String() == other.String()
}

func (u UserName) String() string {
	return string(u)
}
