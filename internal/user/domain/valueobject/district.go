package uservo

import (
	"errors"
	"ifoodish-store/pkg/resperr"
	"net/http"
	"strconv"
)

const (
	MaxDistrictLength = 50
	MinDistrictLength = 5
)

var (
	ErrDistrictMaxLength = errors.New("district should have < " + strconv.Itoa(MaxDistrictLength) + " characteres")
	ErrDistrictMinLength = errors.New("district should have > " + strconv.Itoa(MinDistrictLength) + " characteres")
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
			"O Bairro está muito grande, deve ter menos que"+strconv.Itoa(MaxCityLength)+" digitos",
		)
	}
	if len(value) < MinDistrictLength {
		return "", resperr.WithCodeAndMessage(
			ErrDistrictMinLength,
			http.StatusBadRequest,
			"O Bairro está muito pequeno, deve ter menos que"+strconv.Itoa(MinCityLength)+" digitos",
		)
	}
	return District(value), nil
}
