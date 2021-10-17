package uservo

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
)

var (
	ErrLongitudeInvalidFormat = fmt.Errorf("longitude em formato inválido")
)

type Longitude string

func (s Longitude) Equals(other Longitude) bool {
	return s.String() == other.String()
}

func (s Longitude) String() string {
	return string(s)
}

func NewLongitude(value string) (Longitude, error) {
	if !valid.IsLongitude(value) {
		return "", ErrLongitudeInvalidFormat
	}
	return Longitude(value), nil

}
