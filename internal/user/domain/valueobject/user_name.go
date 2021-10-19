package uservo

import (
	"errors"
	"ifoodish-store/pkg/resperr"
	"net/http"
	"strconv"

	valid "github.com/asaskevich/govalidator"
)

const (
	MinUserNameLength = 3
	MaxUserNameLength = 50
)

var (
	ErrUserNameMaxLength        = errors.New("user name should have < " + strconv.Itoa(MaxStreetLength) + " characteres")
	ErrUserNameMinLength        = errors.New("user name should have > " + strconv.Itoa(MinStreetLength) + " characteres")
	ErrUserNameInvalidCharacter = errors.New("user name should have valid characteres")
)

type UserName string

func NewUserName(value string) (UserName, error) {
	if len(value) > MaxUserNameLength {
		return "", resperr.WithCodeAndMessage(
			ErrUserNameMaxLength,
			http.StatusBadRequest,
			"O nome do usuário está muito grande, deve ter menos que "+strconv.Itoa(MaxStreetLength)+" digitos",
		)
	}
	if len(value) < MinUserNameLength {
		return "", resperr.WithCodeAndMessage(
			ErrUserNameMinLength,
			http.StatusBadRequest,
			"O nome do usuário está muito pequeno, deve ter mais que "+strconv.Itoa(MinStreetLength)+" digitos",
		)
	}

	if !valid.Matches(value, `^[\p{L}\s]*$`) {
		return "", resperr.WithCodeAndMessage(
			ErrUserNameInvalidCharacter,
			http.StatusBadRequest,
			"O nome do usuário é inválido",
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
