package uservo

import (
	"fmt"

	"github.com/carlmjohnson/resperr"

	"net/http"
)

const (
	MaxPhoneLength = 13
	MinPhoneLength = 10
)

var (
	ErrPhoneMaxLength = fmt.Errorf("phone should have max %d characters", MaxPhoneLength)
	ErrPhoneMinLength = fmt.Errorf("phone should have min %d characters", MinPhoneLength)
)

type Phone string

func (p Phone) Equals(other Phone) bool {
	return p.String() == other.String()
}

func (p Phone) String() string {
	return string(p)
}

func NewPhone(value string) (Phone, error) {
	if len(value) > MaxPhoneLength {
		return "", resperr.WithCodeAndMessage(
			ErrPhoneMaxLength,
			http.StatusBadRequest,
			fmt.Sprintf("o telefone deve ter no máximo %d caracteres", MaxPhoneLength),
		)
	}
	if len(value) < MinPhoneLength {
		return "", resperr.WithCodeAndMessage(
			ErrPhoneMinLength,
			http.StatusBadRequest,
			fmt.Sprintf("o telefone deve ter no mínimo %d caracteres", MinPhoneLength),
		)
	}
	return Phone(value), nil
}
