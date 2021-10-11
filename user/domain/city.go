package userdom

import "fmt"

const (
	maxCityLength = 50
)

var (
	ErrCityMaxLength = fmt.Errorf("a cidade deve possuir menos que %d caracteres", maxCityLength)
)

type City string

func (s City) Equals(other City) bool {
	return s.String() == other.String()
}

func (s City) String() string {
	return string(s)
}

func NewCity(value string) (City, error) {
	if len(value) > maxCityLength {
		return "", ErrCityMaxLength
	}
	return City(value), nil
}
