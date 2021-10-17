package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestZipcodeLength(t *testing.T) {
	require := require.New(t)

	zipcode, myError := NewZipcode(strings.Repeat("0", ZipcodeLength+1))
	require.ErrorIs(myError, ErrZipcodeLength)
	require.Len(zipcode, 0)

	zipcode, myError = NewZipcode(strings.Repeat("0", ZipcodeLength-1))
	require.ErrorIs(myError, ErrZipcodeLength)
	require.Len(zipcode, 0)
}

func TestZipcodeCharacteres(t *testing.T) {
	require := require.New(t)

	Zipcode, myError := NewZipcode(strings.Repeat("-", ZipcodeLength))
	require.ErrorIs(myError, ErrZipcodeNotNumeric)
	require.Len(Zipcode, 0)
}

func TestValidZipcode(t *testing.T) {
	require := require.New(t)

	Zipcode, myError := NewZipcode(strings.Repeat("0", ZipcodeLength))
	require.Nil(myError)
	require.NotEmpty(Zipcode)
}

func TestEqualZipcode(t *testing.T) {
	require := require.New(t)

	zipcode, myError := NewZipcode(strings.Repeat("0", ZipcodeLength))
	require.Nil(myError)

	zipcode2, myError2 := NewZipcode(strings.Repeat("0", ZipcodeLength))
	require.Nil(myError2)

	require.True(zipcode.Equals(zipcode2))

}

func TestNotEqualZipcode(t *testing.T) {
	require := require.New(t)

	zipcode, myError := NewZipcode(strings.Repeat("0", ZipcodeLength))
	require.Nil(myError)

	zipcode2, myError2 := NewZipcode(strings.Repeat("1", ZipcodeLength))
	require.Nil(myError2)

	require.False(zipcode.Equals(zipcode2))

}
