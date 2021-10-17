package uservo

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
)

var (
	ErrLatitudeInvalidFormat = fmt.Errorf("latitude inválida")
)

type Latitude string

func (s Latitude) Equals(other Latitude) bool {
	return s.String() == other.String()
}

func (s Latitude) String() string {
	return string(s)
}

func NewLatitude(value string) (Latitude, error) {
	if !valid.IsLatitude(value) {
		return "", ErrLatitudeInvalidFormat
	}
	return Latitude(value), nil

}
