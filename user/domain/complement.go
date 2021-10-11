package userdom

import "fmt"

const (
	maxComplementLength = 300
)

var (
	ErrComplementMaxLength = fmt.Errorf("o complemento deve possuir menos que %d caracteres", maxComplementLength)
)

type Complement string

func (s Complement) Equals(other Complement) bool {
	return s.String() == other.String()
}

func (s Complement) String() string {
	return string(s)
}

func NewComplement(value string) (Complement, error) {
	if len(value) > maxComplementLength {
		return "", ErrComplementMaxLength
	}
	return Complement(value), nil
}
