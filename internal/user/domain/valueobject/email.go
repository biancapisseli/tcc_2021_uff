package uservo

import (
	"errors"
	"ifoodish-store/pkg/resperr"
	"net/http"
	"strconv"

	valid "github.com/asaskevich/govalidator"
)

const (
	MaxEmailLength = 50
)

var (
	ErrEmailMaxLength     = errors.New("email should have < " + strconv.Itoa(MaxEmailLength) + " characteres")
	ErrEmailInvalidFormat = errors.New("email should have a valid format")
)

func (e Email) Equals(other Email) bool {
	return e.String() == other.String()
}

func (e Email) String() string {
	return string(e)
}

type Email string

func NewEmail(value string) (Email, error) {
	if len(value) > MaxEmailLength {
		return "", resperr.WithCodeAndMessage(
			ErrEmailMaxLength,
			http.StatusBadRequest,
			"o Email está muito grande, deve ter menos que "+strconv.Itoa(MaxCityLength)+" digitos",
		)
	}
	if !valid.IsEmail(value) {
		return "", resperr.WithCodeAndMessage(
			ErrEmailInvalidFormat,
			http.StatusBadRequest,
			"o Email está com um formato invalido, siga o padrão 'email@email.com', incluindo .br se necessário",
		)
	}
	return Email(value), nil
}
