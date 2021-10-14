package uservo

import "fmt"

const (
	maxNumberLength = 10
)

var (
	ErrNumberMaxLength = fmt.Errorf("o número deve possuir menos que %d caracteres", maxNumberLength)
)

type Number string

func (s Number) Equals(other Number) bool {
	return s.String() == other.String()
}

func (s Number) String() string {
	return string(s)
}

func NewNumber(value string) (Number, error) {
	if len(value) > maxNumberLength {
		return "", ErrNumberMaxLength
	}
	return Number(value), nil
}
