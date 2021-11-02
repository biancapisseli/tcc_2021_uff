package uservo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddressIDValid(t *testing.T) {
	require := require.New(t)

	id, err := NewAddressID(1)
	require.Nil(err)
	require.NotEmpty(id)

	id, err = NewAddressID(50)
	require.Nil(err)
	require.NotEmpty(id)

	id, err = NewAddressID(100)
	require.Nil(err)
	require.NotEmpty(id)

	id, err = NewAddressID(9999999)
	require.Nil(err)
	require.NotEmpty(id)
}

func TestAddressIDInvalid(t *testing.T) {
	require := require.New(t)

	id, err := NewAddressID(-1)
	require.ErrorIs(err, ErrInvalidAddressID)
	require.Zero(id)

	id, err = NewAddressID(0)
	require.ErrorIs(err, ErrInvalidAddressID)
	require.Zero(id)

	id, err = NewAddressID(-100)
	require.ErrorIs(err, ErrInvalidAddressID)
	require.Zero(id)

}

func TestEqualAddressID(t *testing.T) {
	require := require.New(t)

	AddressID1, err := NewAddressID(50)
	require.Nil(err)

	AddressID2, err := NewAddressID(50)
	require.Nil(err)

	require.True(AddressID1.Equals(AddressID2))

}

func TestNotEqualAddressID(t *testing.T) {
	require := require.New(t)

	AddressID, err := NewAddressID(50)
	require.Nil(err)

	AddressID2, err := NewAddressID(12)
	require.Nil(err)

	require.False(AddressID.Equals(AddressID2))

	AddressID3, err := NewAddressID(500)
	require.Nil(err)

	require.False(AddressID.Equals(AddressID3))

	AddressID4, err := NewAddressID(10)
	require.Nil(err)

	require.False(AddressID.Equals(AddressID4))
}
