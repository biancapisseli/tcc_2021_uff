package uservo

import (
	"fmt"
	"ifoodish-store/pkg/resperr"

	"net/http"
)

const (
	MaxCityLength = 50
	MinCityLength = 2
)

var (
	ErrCityMaxLength = fmt.Errorf("city should have < %d characters", MaxCityLength)
	ErrCityMinLength = fmt.Errorf("city should have > %d characters", MinCityLength)
)

type City string

func (s City) Equals(other City) bool {
	return s.String() == other.String()
}

func (s City) String() string {
	return string(s)
}

func NewCity(value string) (City, error) {
	if len(value) > MaxCityLength {
		return "", resperr.WithCodeAndMessage(
			ErrCityMaxLength,
			http.StatusBadRequest,
			fmt.Sprintf("a cidade deve ter no máximo %d caracteres", MaxCityLength),
		)
	}
	if len(value) < MinCityLength {
		return "", resperr.WithCodeAndMessage(
			ErrCityMinLength,
			http.StatusBadRequest,
			fmt.Sprintf("a cidade deve ter no mínimo %d caracteres", MinCityLength),
		)
	}
	return City(value), nil
}
