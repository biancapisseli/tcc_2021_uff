package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStateMaxLength(t *testing.T) {
	require := require.New(t)

	state, myError := NewState(strings.Repeat("a", maxStateLength+1))
	require.ErrorIs(myError, ErrStateMaxLength)
	require.Len(state, 0)
}

func TestValidState(t *testing.T) {
	require := require.New(t)

	state, myError := NewState(strings.Repeat("a", maxStateLength-1))
	require.Nil(myError)
	require.NotEmpty(state)
}

func TestEqualState(t *testing.T) {
	require := require.New(t)

	state, myError := NewState(strings.Repeat("a", maxStateLength-1))
	require.Nil(myError)

	state2, myError2 := NewState(strings.Repeat("a", maxStateLength-1))
	require.Nil(myError2)

	require.True(state.Equals(state2))

}

func TestNotEqualState(t *testing.T) {
	require := require.New(t)

	state, myError := NewState(strings.Repeat("a", maxStateLength-1))
	require.Nil(myError)

	state2, myError2 := NewState(strings.Repeat("a", maxStateLength-2))
	require.Nil(myError2)

	state3, myError3 := NewState(strings.Repeat("b", maxStateLength-1))
	require.Nil(myError3)

	require.False(state.Equals(state2))
	require.False(state.Equals(state3))

}
