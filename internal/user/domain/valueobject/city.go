package uservo

import "fmt"

const (
	MaxCityLength = 50
	MinCityLength = 2
)

var (
	ErrCityMaxLength = fmt.Errorf("a cidade deve possuir no máximo %d caracteres", MaxCityLength)
	ErrCityMinLength = fmt.Errorf("a cidade deve possuir no mínimo %d caracteres", MaxCityLength)
)

type City string

func (s City) Equals(other City) bool {
	return s.String() == other.String()
}

func (s City) String() string {
	return string(s)
}

func NewCity(value string) (City, error) {
	if len(value) > MaxCityLength {
		return "", ErrCityMaxLength
	}
	if len(value) < MinCityLength {
		return "", ErrCityMinLength
	}
	return City(value), nil
}
