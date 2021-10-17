package uservo

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
)

const (
	MinUserNameLength = 3
	MaxUserNameLength = 50
)

var (
	ErrUserNameMinLength        = fmt.Errorf("nome de usuário deve possui mais de %d caracteres", MinUserNameLength)
	ErrUserNameMaxLength        = fmt.Errorf("nome de usuário deve possui menos de %d caracteres", MaxUserNameLength)
	ErrUserNameInvalidCharacter = fmt.Errorf("nome de usuário deve conter apenas letras e espaços")
)

type UserName string

func NewUserName(value string) (UserName, error) {
	if len(value) > MaxUserNameLength {
		return "", ErrUserNameMaxLength
	}
	if len(value) < MinUserNameLength {
		return "", ErrUserNameMinLength
	}

	if !valid.Matches(value, `^[\p{L}\s]*$`) {
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
