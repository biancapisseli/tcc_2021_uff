package uservo

import (
	"errors"
	"ifoodish-store/pkg/resperr"
	"net/http"
	"strconv"

	valid "github.com/asaskevich/govalidator"
)

const (
	ZipcodeLength = 8
)

var (
	ErrZipcodeLength     = errors.New("zipcode should have " + strconv.Itoa(ZipcodeLength) + " characteres")
	ErrZipcodeNotNumeric = errors.New("zipcode should be valid ")
)

type Zipcode string

func (s Zipcode) Equals(other Zipcode) bool {
	return s.String() == other.String()
}

func (s Zipcode) String() string {
	return string(s)
}

func NewZipcode(value string) (Zipcode, error) {
	if len(value) != ZipcodeLength {
		return "", resperr.WithCodeAndMessage(
			ErrZipcodeLength,
			http.StatusBadRequest,
			"O CEP deve ter "+strconv.Itoa(ZipcodeLength)+" dígitos",
		)
	}
	if !valid.IsNumeric(value) {
		return "", resperr.WithCodeAndMessage(
			ErrZipcodeNotNumeric,
			http.StatusBadRequest,
			"O CEP deve ser numérico",
		)
	}
	return Zipcode(value), nil
}
