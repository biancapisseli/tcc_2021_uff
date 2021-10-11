package userdom

import "fmt"

const (
	maxStreetLength = 50
	minStreetLength = 5
)

var (
	ErrStreetMaxLength = fmt.Errorf("a rua deve possuir menos que %d caracteres", maxStreetLength)
	ErrStreetMinLength = fmt.Errorf("a rua deve possuir mais que %d caracteres", minStreetLength)
)

type Street string

func (s Street) Equals(other Street) bool {
	return s.String() == other.String()
}

func (s Street) String() string {
	return string(s)
}

func NewStreet(value string) (Street, error) {
	if len(value) > maxStreetLength {
		return "", ErrStreetMaxLength
	}
	if len(value) < minStreetLength {
		return "", ErrStreetMinLength
	}
	return Street(value), nil
}
