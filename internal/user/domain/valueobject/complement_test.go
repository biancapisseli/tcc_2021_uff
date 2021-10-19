package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvalidComplement(t *testing.T) {
	require := require.New(t)

	complement, err := NewComplement(strings.Repeat("a", MaxComplementLength+1))
	require.ErrorIs(err, ErrComplementMaxLength)

	require.Len(complement, 0)

}

func TestValidComplement(t *testing.T) {
	require := require.New(t)

	complement, err := NewComplement(strings.Repeat("a", MaxComplementLength))
	require.Nil(err)

	require.NotEmpty(complement)

	complement, err = NewComplement("")
	require.Nil(err)

	require.Len(complement, 0)
}

func TestEqualComplement(t *testing.T) {
	require := require.New(t)

	complement, err := NewComplement(strings.Repeat("a", MaxComplementLength))
	require.Nil(err)

	complement2, err2 := NewComplement(strings.Repeat("a", MaxComplementLength))
	require.Nil(err2)

	require.True(complement.Equals(complement2))

}

func TestNotEqualComplement(t *testing.T) {
	require := require.New(t)

	complement, err := NewComplement(strings.Repeat("a", MaxComplementLength))
	require.Nil(err)

	complement2, err2 := NewComplement(strings.Repeat("b", MaxComplementLength))
	require.Nil(err2)

	complement3, err3 := NewComplement(strings.Repeat("a", MaxComplementLength-1))
	require.Nil(err3)

	require.False(complement.Equals(complement3))
	require.False(complement.Equals(complement2))

}
