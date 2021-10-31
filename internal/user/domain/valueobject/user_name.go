package uservo

import (
	"errors"
	"fmt"
	"ifoodish-store/pkg/resperr"
	"net/http"
	valid "github.com/asaskevich/govalidator"
)

const (
	MinUserNameLength = 3
	MaxUserNameLength = 50
)

var (
	ErrUserNameMaxLength        = fmt.Errorf("user name should have < %d characters", MaxStreetLength)
	ErrUserNameMinLength        = fmt.Errorf("user name should have > %d characters", MinStreetLength)
	ErrUserNameInvalidCharacter = errors.New("user name should have only letters and spaces")
)

type UserName string

func NewUserName(value string) (UserName, error) {
	if len(value) > MaxUserNameLength {
		return "", resperr.WithCodeAndMessage(
			ErrUserNameMaxLength,
			http.StatusBadRequest,
			fmt.Sprintf("o nome do usuário deve ter no máximo %d caracteres", MaxStreetLength),
		)
	}
	if len(value) < MinUserNameLength {
		return "", resperr.WithCodeAndMessage(
			ErrUserNameMinLength,
			http.StatusBadRequest,
			fmt.Sprintf("o nome do usuário deve ter no mínimo %d caracteres", MinStreetLength),
		)
	}

	if !valid.Matches(value, `^[\p{L}\s]*$`) {
		return "", resperr.WithCodeAndMessage(
			ErrUserNameInvalidCharacter,
			http.StatusBadRequest,
			"o nome do usuário é inválido",
		)
	}
	return UserName(value), nil
}

func (u UserName) Equals(other UserName) bool {
	return u.String() == other.String()
}

func (u UserName) String() string {
	return string(u)
}
