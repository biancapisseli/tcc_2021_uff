package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvalidState(t *testing.T) {
	require := require.New(t)

	state, err := NewState(strings.Repeat("a", MaxStateLength+1))
	require.ErrorIs(err, ErrStateMaxLength)
	require.Len(state, 0)

	state, err = NewState(strings.Repeat("a", MinStateLength-1))
	require.ErrorIs(err, ErrStateMinLength)
	require.Len(state, 0)
}

func TestValidState(t *testing.T) {
	require := require.New(t)

	state, err := NewState(strings.Repeat("a", MaxStateLength-1))
	require.Nil(err)
	require.NotEmpty(state)
}

func TestEqualState(t *testing.T) {
	require := require.New(t)

	state, err := NewState(strings.Repeat("a", MaxStateLength-1))
	require.Nil(err)

	state2, err2 := NewState(strings.Repeat("a", MaxStateLength-1))
	require.Nil(err2)

	require.True(state.Equals(state2))

}

func TestNotEqualState(t *testing.T) {
	require := require.New(t)

	state, err := NewState(strings.Repeat("a", MaxStateLength-1))
	require.Nil(err)

	state2, err2 := NewState(strings.Repeat("a", MaxStateLength-2))
	require.Nil(err2)

	state3, err3 := NewState(strings.Repeat("b", MaxStateLength-1))
	require.Nil(err3)

	require.False(state.Equals(state2))
	require.False(state.Equals(state3))

}
