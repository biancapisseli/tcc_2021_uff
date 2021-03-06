package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvalidStreet(t *testing.T) {
	require := require.New(t)

	street, myError := NewStreet(strings.Repeat("a", MaxStreetLength+1))
	require.ErrorIs(myError, ErrStreetMaxLength)
	require.Len(street, 0)

	street, myError = NewStreet(strings.Repeat("a", MinStreetLength-1))
	require.ErrorIs(myError, ErrStreetMinLength)
	require.Len(street, 0)
}

func TestValidStreet(t *testing.T) {
	require := require.New(t)

	street, myError := NewStreet(strings.Repeat("a", MaxStreetLength-1))
	require.Nil(myError)
	require.NotEmpty(street)
}

func TestEqualStreet(t *testing.T) {
	require := require.New(t)

	street, myError := NewStreet(strings.Repeat("a", MaxStreetLength-1))
	require.Nil(myError)

	street2, myError2 := NewStreet(strings.Repeat("a", MaxStreetLength-1))
	require.Nil(myError2)

	require.True(street.Equals(street2))

}

func TestNotEqualStreet(t *testing.T) {
	require := require.New(t)

	street, myError := NewStreet(strings.Repeat("a", MaxStreetLength-1))
	require.Nil(myError)

	street2, myError2 := NewStreet(strings.Repeat("a", MaxStreetLength))
	require.Nil(myError2)

	street3, myError3 := NewStreet(strings.Repeat("b", MaxStreetLength-1))
	require.Nil(myError3)

	require.False(street.Equals(street2))
	require.False(street.Equals(street3))

}
