package uservo

import "fmt"

const (
	maxStateLength = 50
)

var (
	ErrStateMaxLength = fmt.Errorf("o estado deve possuir menos que %d caracteres", maxStateLength)
)

type State string

func (s State) Equals(other State) bool {
	return s.String() == other.String()
}

func (s State) String() string {
	return string(s)
}

func NewState(value string) (State, error) {
	if len(value) > maxStateLength {
		return "", ErrStateMaxLength
	}
	return State(value), nil
}
