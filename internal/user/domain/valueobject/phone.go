package uservo

import "fmt"

const (
	MaxPhoneLength = 11
	MinPhoneLength = 10
)

var (
	ErrPhoneMaxLength = fmt.Errorf("o telefone deve possuir no máximo %d caracteres", MaxPhoneLength)
	ErrPhoneMinLength = fmt.Errorf("o telefone deve possuir mais que %d caracteres", MinPhoneLength)
)

type Phone string

func (p Phone) Equals(other Phone) bool {
	return p.String() == other.String()
}

func (p Phone) String() string {
	return string(p)
}

func NewPhone(value string) (Phone, error) {
	if len(value) > MaxPhoneLength {
		return "", ErrPhoneMaxLength
	}
	if len(value) < MinPhoneLength {
		return "", ErrPhoneMinLength
	}
	return Phone(value), nil
}
