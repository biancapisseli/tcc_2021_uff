package userdom

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
)

const (
	maxEmailLength = 50
)

var (
	ErrEmailMaxLength     = fmt.Errorf("email deve possuir menos de %d caracteres", maxEmailLength)
	ErrEmailInvalidFormat = fmt.Errorf("utilize o padrão para email: example@mail.com")
)

type Email string

func NewEmail(value string) (Email, error) {
	if len(value) > maxEmailLength {
		return "", ErrEmailMaxLength
	}
	if !valid.IsEmail(value) {
		return "", ErrEmailInvalidFormat
	}
	return Email(value), nil
}

func (e Email) Equals(other Email) bool {
	return e.String() == other.String()
}

func (e Email) String() string {
	return string(e)
}
