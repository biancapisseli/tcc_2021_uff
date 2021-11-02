package uservo

import (
	"fmt"

	"github.com/carlmjohnson/resperr"

	"net/http"
)

const (
	MaxDistrictLength = 50
	MinDistrictLength = 5
)

var (
	ErrDistrictMaxLength = fmt.Errorf("district should have max %d characters", MaxDistrictLength)
	ErrDistrictMinLength = fmt.Errorf("district should have min %d characters", MinDistrictLength)
)

type District string

func (s District) Equals(other District) bool {
	return s.String() == other.String()
}

func (s District) String() string {
	return string(s)
}

func NewDistrict(value string) (District, error) {
	if len(value) > MaxDistrictLength {
		return "", resperr.WithCodeAndMessage(
			ErrDistrictMaxLength,
			http.StatusBadRequest,
			fmt.Sprintf("o bairro deve ter no máximo %d caracteres", MaxCityLength),
		)
	}
	if len(value) < MinDistrictLength {
		return "", resperr.WithCodeAndMessage(
			ErrDistrictMinLength,
			http.StatusBadRequest,
			fmt.Sprintf("o bairro deve ter no máximo %d caracteres", MinCityLength),
		)
	}
	return District(value), nil
}
