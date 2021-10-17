package uservo

import "fmt"

const (
	MaxDistrictLength = 50
	MinDistrictLength = 5
)

var (
	ErrDistrictMaxLength = fmt.Errorf("o bairro deve possuir no máximo %d caracteres", MaxDistrictLength)
	ErrDistrictMinLength = fmt.Errorf("o bairro deve possuir mais que %d caracteres", MinDistrictLength)
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
		return "", ErrDistrictMaxLength
	}
	if len(value) < MinDistrictLength {
		return "", ErrDistrictMinLength
	}
	return District(value), nil
}
