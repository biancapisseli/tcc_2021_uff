package uservo

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
)

const (
	ZipcodeLength = 8
)

var (
	ErrZipcodeLength     = fmt.Errorf("o CEP deve possuir %d digitos", ZipcodeLength)
	ErrZipcodeNotNumeric = fmt.Errorf("o CEP deve possuir apenas numeros")
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
		return "", ErrZipcodeLength
	}
	if !valid.IsNumeric(value) {
		return "", ErrZipcodeNotNumeric
	}
	return Zipcode(value), nil
}
