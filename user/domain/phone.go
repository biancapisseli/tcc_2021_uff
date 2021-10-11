package userdom

import "fmt"

const (
	maxPhoneLength = 11
	minPhoneLength = 10
)

var (
	ErrPhoneMaxLength = fmt.Errorf("o telefone deve possuir menos que %d caracteres", maxPhoneLength)
	ErrPhoneMinLength = fmt.Errorf("o telefone deve possuir mais que %d caracteres", minPhoneLength)
)

type Phone string

func (p Phone) Equals(other Phone) bool {
	return p.String() == other.String()
}

func (p Phone) String() string {
	return string(p)
}

func NewPhone(value string) (Phone, error) {
	if len(value) > maxPhoneLength {
		return "", ErrPhoneMaxLength
	}
	if len(value) < minPhoneLength {
		return "", ErrPhoneMinLength
	}
	return Phone(value), nil
}
