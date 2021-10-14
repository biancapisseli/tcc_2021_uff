package uservo

import "fmt"

const (
	maxDistrictLength = 50
	minDistrictLength = 5
)

var (
	ErrDistrictMaxLength = fmt.Errorf("o bairro deve possuir menos que %d caracteres", maxDistrictLength)
	ErrDistrictMinLength = fmt.Errorf("o bairro deve possuir mais que %d caracteres", minDistrictLength)
)

type District string

func (s District) Equals(other District) bool {
	return s.String() == other.String()
}

func (s District) String() string {
	return string(s)
}

func NewDistrict(value string) (District, error) {
	if len(value) > maxDistrictLength {
		return "", ErrDistrictMaxLength
	}
	if len(value) < minDistrictLength {
		return "", ErrDistrictMinLength
	}
	return District(value), nil
}
