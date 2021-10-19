package uservo

import (
	"errors"
	"ifoodish-store/pkg/resperr"
	"net/http"
	"strconv"
)

const (
	MaxComplementLength = 300
)

var (
	ErrComplementMaxLength = errors.New("complement should have < " + strconv.Itoa(MaxComplementLength) + " characteres")
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
			"O complemento está muito grande, deve ter menos que "+strconv.Itoa(MaxComplementLength)+" digitos",
		)
	}
	return Complement(value), nil
}
