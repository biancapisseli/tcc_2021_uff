package uservo

import (
	"fmt"
	"ifoodish-store/pkg/resperr"

	"net/http"
)

const (
	MaxStateLength = 50
	MinStateLength = 2
)

var (
	ErrStateMaxLength = fmt.Errorf("state should have < %d characters", MaxStateLength)
	ErrStateMinLength = fmt.Errorf("state should have > %d characters", MinStateLength)
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
			fmt.Sprintf("o estado deve ter no máximo %d caracteres", MaxStateLength),
		)
	}
	if len(value) < MinStateLength {
		return "", resperr.WithCodeAndMessage(
			ErrStateMinLength,
			http.StatusBadRequest,
			fmt.Sprintf("o estado deve ter no mínimo %d caracteres", MinStateLength),
		)
	}
	return State(value), nil
}
