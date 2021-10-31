package uservo

import (
	"github.com/carlmjohnson/resperr"

	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUserIDInvalid(t *testing.T) {
	require := require.New(t)

	id, err := NewUserID("")
	require.NotNil(err)
	require.Equal(http.StatusBadRequest, resperr.StatusCode(err))
	require.Equal(id, UserID(uuid.Nil))
}

func TestUserIDValid(t *testing.T) {
	require := require.New(t)

	id, err := NewUserID("123e4567-e89b-12d3-a456-426614174000")
	require.Nil(err)
	require.Equal("123e4567-e89b-12d3-a456-426614174000", id.String())

	genaratedId := GenerateNewUserID()
	id, err = NewUserID(genaratedId.String())
	require.Nil(err)
	require.Equal(genaratedId, id)

}

func TestUserIDString(t *testing.T) {
	require := require.New(t)

	id, err := NewUserID("123e4567-e89b-12d3-a456-426614174000")
	require.Nil(err)
	require.Equal("123e4567-e89b-12d3-a456-426614174000", id.String())
}

func TestEqualUserID(t *testing.T) {
	require := require.New(t)

	userID, err := NewUserID("123e4567-e89b-12d3-a456-426614174000")
	require.Nil(err)

	userID2, err2 := NewUserID("123e4567-e89b-12d3-a456-426614174000")
	require.Nil(err2)

	require.True(userID.Equals(userID2))
}

func TestNotEqualUserID(t *testing.T) {
	require := require.New(t)

	userID, err := NewUserID("123e4567-e89b-12d3-a456-426614174000")
	require.Nil(err)

	userID2, err2 := NewUserID("f2712cc3-6cd4-4690-b995-68d4ab51e908")
	require.Nil(err2)

	userID3, err3 := NewUserID("da548ecb-688f-4685-97dc-622973540288")
	require.Nil(err3)

	require.False(userID.Equals(userID2))
	require.False(userID.Equals(userID3))

}
