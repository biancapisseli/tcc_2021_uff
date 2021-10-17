package uservo

import "fmt"

const (
	MaxAddressNumberLength = 10
	MinAddressNumberLength = 1
)

var (
	ErrAddressNumberMaxLength = fmt.Errorf("o número deve possuir no máximo %d caracteres", MaxAddressNumberLength)
	ErrAddressNumberMinLength = fmt.Errorf("o número deve possuir no mínimo %d caracteres", MinAddressNumberLength)
)

type Number string

func (s Number) Equals(other Number) bool {
	return s.String() == other.String()
}

func (s Number) String() string {
	return string(s)
}

func NewAddressNumber(value string) (Number, error) {
	if len(value) > MaxAddressNumberLength {
		return "", ErrAddressNumberMaxLength
	}
	if len(value) < MinAddressNumberLength {
		return "", ErrAddressNumberMinLength
	}

	return Number(value), nil
}
