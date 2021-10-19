package uservo

import (
	"errors"
	"ifoodish-store/pkg/resperr"
	"net/http"
	"strconv"
)

const (
	MaxCityLength = 50
	MinCityLength = 2
)

var (
	ErrCityMaxLength = errors.New("city should have < " + strconv.Itoa(MaxCityLength) + " characteres")
	ErrCityMinLength = errors.New("city should have > " + strconv.Itoa(MinCityLength) + " characteres")
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
			"A Cidade está muito grande, deve ter menos que "+strconv.Itoa(MaxCityLength)+" digitos",
		)
	}
	if len(value) < MinCityLength {
		return "", resperr.WithCodeAndMessage(
			ErrCityMinLength,
			http.StatusBadRequest,
			"A Cidade está muito pequeno, deve ter mais que "+strconv.Itoa(MinCityLength)+" digitos",
		)
	}
	return City(value), nil
}
