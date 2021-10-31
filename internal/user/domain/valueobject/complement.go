package uservo

import (
	"fmt"
	"ifoodish-store/pkg/resperr"
	"net/http"
)

const (
	MaxComplementLength = 300
)

var (
	ErrComplementMaxLength = fmt.Errorf("complement should have < %d characters", MaxComplementLength)
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
		return "", resperr.WithCodeAndMessage(
			ErrComplementMaxLength,
			http.StatusBadRequest,
			fmt.Sprintf("o complemento deve ter no máximo %d caracteres", MaxComplementLength),
		)
	}
	return Complement(value), nil
}
