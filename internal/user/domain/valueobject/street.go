package uservo

import "fmt"

const (
	MaxStreetLength = 50
	MinStreetLength = 5
)

var (
	ErrStreetMaxLength = fmt.Errorf("a rua deve possuir no máximo %d caracteres", MaxStreetLength)
	ErrStreetMinLength = fmt.Errorf("a rua deve possuir mais que %d caracteres", MinStreetLength)
)

type Street string

func (s Street) Equals(other Street) bool {
	return s.String() == other.String()
}

func (s Street) String() string {
	return string(s)
}

func NewStreet(value string) (Street, error) {
	if len(value) > MaxStreetLength {
		return "", ErrStreetMaxLength
	}
	if len(value) < MinStreetLength {
		return "", ErrStreetMinLength
	}
	return Street(value), nil
}
