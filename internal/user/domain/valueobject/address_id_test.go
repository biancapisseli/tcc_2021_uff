package uservo

import (
	"testing"
	"github.com/stretchr/testify/require"
)

const (
	validID   = 50
	invalidID = -1
)

func TestAddressIDValid(t *testing.T) {
	require := require.New(t)

	id, err := NewAddressID(validID)
	require.Nil(err)
	require.NotEmpty(id)
}

func TestAddressIDInvalid(t *testing.T) {
	require := require.New(t)

	id, err := NewAddressID(invalidID)
	require.ErrorIs(err, ErrInvalidAddressID)
	require.Zero(id)
}

func TestEqualAddressID(t *testing.T) {
	require := require.New(t)

	AddressID, err := NewAddressID(validID)
	require.Nil(err)

	AddressIDEqual, errEqual := NewAddressID(validID)
	require.Nil(errEqual)

	AddressIDInt, errInt := NewAddressID(50)
	require.Nil(errInt)

	require.True(AddressID.Equals(AddressIDEqual))
	require.True(AddressID.Equals(AddressIDInt))
}

func TestNotEqualAddressID(t *testing.T) {
	require := require.New(t)

	AddressID, err := NewAddressID(validID)
	require.Nil(err)

	AddressIDEqual, errEqual := NewAddressID(12)
	require.Nil(errEqual)

	require.False(AddressID.Equals(AddressIDEqual))

}
