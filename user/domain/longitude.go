package userdom

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
)

var (
	ErrLongitudeFormat = fmt.Errorf("longitude inválida")
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
		return "", ErrLongitudeFormat
	}
	return Longitude(value), nil

}
