package uservo

import (
	"errors"
	"ifoodish-store/pkg/resperr"
	"net/http"
	"strconv"
)

const (
	MaxStateLength = 50
	MinStateLength = 2
)

var (
	ErrStateMaxLength = errors.New("state should have < " + strconv.Itoa(MaxStateLength) + " characteres")
	ErrStateMinLength = errors.New("state should have > " + strconv.Itoa(MinStateLength) + " characteres")
)

type State string

func (s State) Equals(other State) bool {
	return s.String() == other.String()
}

func (s State) String() string {
	return string(s)
}

func NewState(value string) (State, error) {
	if len(value) > MaxStateLength {
		return "", resperr.WithCodeAndMessage(
			ErrStateMaxLength,
			http.StatusBadRequest,
			"O estado está muito grande, deve ter menos que"+strconv.Itoa(MaxStateLength)+" digitos",
		)
	}
	if len(value) < MinStateLength {
		return "", resperr.WithCodeAndMessage(
			ErrStateMinLength,
			http.StatusBadRequest,
			"O estado está muito pequeno, deve ter mais que"+strconv.Itoa(MinStateLength)+" digitos",
		)
	}
	return State(value), nil
}
