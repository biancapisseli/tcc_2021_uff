package uservo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	validID      = 50
	validInt     = 50
	validIDEqual = 12
	invalidID    = -1
)

func TestAddressIDValid(t *testing.T) {
	require := require.New(t)

	id, myError := NewAddressID(validID)
	require.Nil(myError)
	require.NotEmpty(id)
}

func TestAddressIDInvalid(t *testing.T) {
	require := require.New(t)

	id, myError := NewAddressID(invalidID)
	require.ErrorIs(myError, ErrInvalidAddressID)
	require.ElementsMatch(id, Invalid)
}

func TestEqualAddressID(t *testing.T) {
	require := require.New(t)

	AddressID, myError := NewAddressID(validID)
	require.Nil(myError)

	AddressIDEqual, myErrorEqual := NewAddressID(validID)
	require.Nil(myErrorEqual)

	AddressIDInt, myErrorInt := NewAddressID(validInt)
	require.Nil(myErrorInt)

	require.True(AddressID.Equals(AddressIDEqual))
	require.True(AddressID.Equals(AddressIDInt))
}

func TestNotEqualAddressID(t *testing.T) {
	require := require.New(t)

	AddressID, myError := NewAddressID(validID)
	require.Nil(myError)

	AddressIDEqual, myErrorEqual := NewAddressID(validIDEqual)
	require.Nil(myErrorEqual)

	require.False(AddressID.Equals(AddressIDEqual))

}
