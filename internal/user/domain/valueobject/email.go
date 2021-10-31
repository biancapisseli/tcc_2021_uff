package uservo

import (
	"fmt"

	"github.com/carlmjohnson/resperr"

	"net/http"

	valid "github.com/asaskevich/govalidator"
)

const (
	MaxEmailLength = 50
)

var (
	ErrEmailMaxLength     = fmt.Errorf("email should have < %d characters", MaxEmailLength)
	ErrEmailInvalidFormat = fmt.Errorf("email should have a valid format")
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
			fmt.Sprintf("o email deve ter no máximo %d caracteres", MaxCityLength),
		)
	}
	if !valid.IsEmail(value) {
		return "", resperr.WithCodeAndMessage(
			ErrEmailInvalidFormat,
			http.StatusBadRequest,
			`o email deve ter o formato 'joao.silva@email.com', incluindo .br se necessário`,
		)
	}
	return Email(value), nil
}
