package userdom

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStreetMaxLength(t *testing.T) {
	require := require.New(t)

	street, myError := NewStreet(strings.Repeat("a", maxStreetLength+1))
	require.ErrorIs(myError, ErrStreetMaxLength)
	require.Len(street, 0)
}

func TestStreetMinLength(t *testing.T) {
	require := require.New(t)

	street, myError := NewStreet(strings.Repeat("a", minStreetLength-1))
	require.ErrorIs(myError, ErrStreetMinLength)
	require.Len(street, 0)
}

func TestValidStreet(t *testing.T) {
	require := require.New(t)

	street, myError := NewStreet(strings.Repeat("a", maxStreetLength-1))
	require.Nil(myError)
	require.NotEmpty(street)
}

func TestEqualStreet(t *testing.T) {
	require := require.New(t)

	street, myError := NewStreet(strings.Repeat("a", maxStreetLength-1))
	require.Nil(myError)

	street2, myError2 := NewStreet(strings.Repeat("a", maxStreetLength-1))
	require.Nil(myError2)

	require.True(street.Equals(street2))

}

func TestNotEqualStreet(t *testing.T) {
	require := require.New(t)

	street, myError := NewStreet(strings.Repeat("a", maxStreetLength-1))
	require.Nil(myError)

	street2, myError2 := NewStreet(strings.Repeat("a", maxStreetLength))
	require.Nil(myError2)

	street3, myError3 := NewStreet(strings.Repeat("b", maxStreetLength-1))
	require.Nil(myError3)

	require.False(street.Equals(street2))
	require.False(street.Equals(street3))

}
