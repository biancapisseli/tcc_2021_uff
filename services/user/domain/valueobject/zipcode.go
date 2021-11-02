package uservo

import (
	"errors"
	"fmt"

	"github.com/carlmjohnson/resperr"

	"net/http"

	valid "github.com/asaskevich/govalidator"
)

const (
	ZipcodeLength = 8
)

var (
	ErrZipcodeLength     = fmt.Errorf("zipcode should have %d characters", ZipcodeLength)
	ErrZipcodeNotNumeric = errors.New("zipcode should be numeric")
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
			fmt.Sprintf("o CEP deve ter %d dígitos", ZipcodeLength),
		)
	}
	if !valid.IsNumeric(value) {
		return "", resperr.WithCodeAndMessage(
			ErrZipcodeNotNumeric,
			http.StatusBadRequest,
			"o CEP deve ser numérico",
		)
	}
	return Zipcode(value), nil
}
