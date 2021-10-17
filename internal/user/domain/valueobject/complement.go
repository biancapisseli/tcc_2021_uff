package uservo

import "fmt"

const (
	MaxComplementLength = 300
)

var (
	ErrComplementMaxLength = fmt.Errorf("o complemento deve possuir no máximo %d caracteres", MaxComplementLength)
)

type Complement string

func (s Complement) Equals(other Complement) bool {
	return s.String() == other.String()
}

func (s Complement) String() string {
	return string(s)
}

func NewComplement(value string) (Complement, error) {
	if len(value) > MaxComplementLength {
		return "", ErrComplementMaxLength
	}
	return Complement(value), nil
}
