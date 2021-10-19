package uservo

import (
	"errors"
	"ifoodish-store/pkg/resperr"
	"net/http"
	"strconv"
)

const (
	MaxStreetLength = 50
	MinStreetLength = 5
)

var (
	ErrStreetMaxLength = errors.New("street should have < " + strconv.Itoa(MaxStreetLength) + "characteres")
	ErrStreetMinLength = errors.New("street should have > " + strconv.Itoa(MinStreetLength) + "characteres")
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
			"A rua está muito grande, deve ter menos que "+strconv.Itoa(MaxStreetLength)+" digitos",
		)
	}
	if len(value) < MinStreetLength {
		return "", resperr.WithCodeAndMessage(
			ErrStreetMinLength,
			http.StatusBadRequest,
			"A rua está muito pequena, deve ter mais que "+strconv.Itoa(MinStreetLength)+" digitos",
		)
	}
	return Street(value), nil
}
