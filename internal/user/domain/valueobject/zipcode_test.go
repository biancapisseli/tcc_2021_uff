package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvalidZipcode(t *testing.T) {
	require := require.New(t)

	zipcode, err := NewZipcode(strings.Repeat("0", ZipcodeLength+1))
	require.ErrorIs(err, ErrZipcodeLength)
	require.Len(zipcode, 0)

	zipcode, err = NewZipcode(strings.Repeat("0", ZipcodeLength-1))
	require.ErrorIs(err, ErrZipcodeLength)
	require.Len(zipcode, 0)

	Zipcode, err := NewZipcode(strings.Repeat("-", ZipcodeLength))
	require.ErrorIs(err, ErrZipcodeNotNumeric)
	require.Len(Zipcode, 0)
}

func TestValidZipcode(t *testing.T) {
	require := require.New(t)

	Zipcode, err := NewZipcode(strings.Repeat("0", ZipcodeLength))
	require.Nil(err)
	require.NotEmpty(Zipcode)
}

func TestEqualZipcode(t *testing.T) {
	require := require.New(t)

	zipcode, err := NewZipcode(strings.Repeat("0", ZipcodeLength))
	require.Nil(err)

	zipcode2, err2 := NewZipcode(strings.Repeat("0", ZipcodeLength))
	require.Nil(err2)

	require.True(zipcode.Equals(zipcode2))

}

func TestNotEqualZipcode(t *testing.T) {
	require := require.New(t)

	zipcode, err := NewZipcode(strings.Repeat("0", ZipcodeLength))
	require.Nil(err)

	zipcode2, err2 := NewZipcode(strings.Repeat("1", ZipcodeLength))
	require.Nil(err2)

	require.False(zipcode.Equals(zipcode2))

}
