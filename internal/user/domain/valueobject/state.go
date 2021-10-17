package uservo

import "fmt"

const (
	MaxStateLength = 50
	MinStateLength = 2
)

var (
	ErrStateMaxLength = fmt.Errorf("o estado deve possuir no máximo %d caracteres", MaxStateLength)
	ErrStateMinLength = fmt.Errorf("o estado deve possuir no mínimo %d caracteres", MinStateLength)
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
		return "", ErrStateMaxLength
	}
	if len(value) < MinStateLength {
		return "", ErrStateMinLength
	}
	return State(value), nil
}
