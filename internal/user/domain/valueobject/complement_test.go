package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComplementMaxLength(t *testing.T) {
	require := require.New(t)

	complement, myError := NewComplement(strings.Repeat("a", MaxComplementLength+1))
	require.ErrorIs(myError, ErrComplementMaxLength)

	require.Len(complement, 0)
}

func TestValidComplement(t *testing.T) {
	require := require.New(t)

	complement, myError := NewComplement(strings.Repeat("a", MaxComplementLength))
	require.Nil(myError)

	require.NotEmpty(complement)
}

func TestEqualComplement(t *testing.T) {
	require := require.New(t)

	complement, myError := NewComplement(strings.Repeat("a", MaxComplementLength))
	require.Nil(myError)

	complement2, myError2 := NewComplement(strings.Repeat("a", MaxComplementLength))
	require.Nil(myError2)

	require.True(complement.Equals(complement2))

}

func TestNotEqualComplement(t *testing.T) {
	require := require.New(t)

	complement, myError := NewComplement(strings.Repeat("a", MaxComplementLength))
	require.Nil(myError)

	complement2, myError2 := NewComplement(strings.Repeat("b", MaxComplementLength))
	require.Nil(myError2)

	complement3, myError3 := NewComplement(strings.Repeat("a", MaxComplementLength-1))
	require.Nil(myError3)

	require.False(complement.Equals(complement3))
	require.False(complement.Equals(complement2))

}
