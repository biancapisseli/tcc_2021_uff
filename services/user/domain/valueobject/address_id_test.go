package uservo

import (
	"github.com/carlmjohnson/resperr"

	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAddressIDInvalid(t *testing.T) {
	require := require.New(t)

	id, err := NewAddressID("")
	require.NotNil(err)
	require.Equal(http.StatusBadRequest, resperr.StatusCode(err))
	require.Equal(id, AddressID(uuid.Nil))
}

func TestAddressIDValid(t *testing.T) {
	require := require.New(t)

	id, err := NewAddressID("123e4567-e89b-12d3-a456-426614174000")
	require.Nil(err)
	require.Equal("123e4567-e89b-12d3-a456-426614174000", id.String())

	genaratedId := GenerateNewAddressID()
	id, err = NewAddressID(genaratedId.String())
	require.Nil(err)
	require.Equal(genaratedId, id)

}

func TestAddressIDString(t *testing.T) {
	require := require.New(t)

	id, err := NewAddressID("123e4567-e89b-12d3-a456-426614174000")
	require.Nil(err)
	require.Equal("123e4567-e89b-12d3-a456-426614174000", id.String())
}

func TestEqualAddressID(t *testing.T) {
	require := require.New(t)

	userID, err := NewAddressID("123e4567-e89b-12d3-a456-426614174000")
	require.Nil(err)

	userID2, err2 := NewAddressID("123e4567-e89b-12d3-a456-426614174000")
	require.Nil(err2)

	require.True(userID.Equals(userID2))
}

func TestNotEqualAddressID(t *testing.T) {
	require := require.New(t)

	userID, err := NewAddressID("123e4567-e89b-12d3-a456-426614174000")
	require.Nil(err)

	userID2, err2 := NewAddressID("f2712cc3-6cd4-4690-b995-68d4ab51e908")
	require.Nil(err2)

	userID3, err3 := NewAddressID("da548ecb-688f-4685-97dc-622973540288")
	require.Nil(err3)

	require.False(userID.Equals(userID2))
	require.False(userID.Equals(userID3))

}
