package uservo

import (
	"fmt"

	"github.com/carlmjohnson/resperr"

	"net/http"
)

const (
	MaxStreetLength = 50
	MinStreetLength = 5
)

var (
	ErrStreetMaxLength = fmt.Errorf("street should have < %d characters", MaxStreetLength)
	ErrStreetMinLength = fmt.Errorf("street should have > %d characters", MinStreetLength)
)

type Street string

func (s Street) Equals(other Street) bool {
	return s.String() == other.String()
}

func (s Street) String() string {
	return string(s)
}

func NewStreet(value string) (Street, error) {
	if len(value) > MaxStreetLength {
		return "", resperr.WithCodeAndMessage(
			ErrStreetMaxLength,
			http.StatusBadRequest,
			fmt.Sprintf("a rua deve ter no máximo %d caracteres", MaxStreetLength),
		)
	}
	if len(value) < MinStreetLength {
		return "", resperr.WithCodeAndMessage(
			ErrStreetMinLength,
			http.StatusBadRequest,
			fmt.Sprintf("a rua deve ter no mínimo %d caracteres", MinStreetLength),
		)
	}
	return Street(value), nil
}
