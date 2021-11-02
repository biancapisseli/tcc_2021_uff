package uservo

import (
	"errors"

	"github.com/carlmjohnson/resperr"

	"net/http"

	valid "github.com/asaskevich/govalidator"
)

var (
	ErrLongitudeInvalidFormat = errors.New("longitude should have a valid format")
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
		return "", resperr.WithCodeAndMessage(
			ErrLongitudeInvalidFormat,
			http.StatusBadRequest,
			`a longitude deve estar no formato "X.XXXX" ou "-X.XXXX"`,
		)
	}
	return Longitude(value), nil

}
