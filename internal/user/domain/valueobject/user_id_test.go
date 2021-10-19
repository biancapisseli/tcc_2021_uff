package uservo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserIDInvalid(t *testing.T) {
	require := require.New(t)

	id, err := NewUserID(-1)
	require.ErrorIs(err, ErrInvalidUserID)
	require.ElementsMatch(id, 0)
}

func TestUserIDValid(t *testing.T) {
	require := require.New(t)

	id, err := NewUserID(5)
	require.Nil(err)
	require.NotEmpty(id)
}

func TestEqualUserID(t *testing.T) {
	require := require.New(t)

	userID, err := NewUserID(11)
	require.Nil(err)

	userID2, err2 := NewUserID(11)
	require.Nil(err2)

	require.True(userID.Equals(userID2))
}

func TestNotEqualUserID(t *testing.T) {
	require := require.New(t)

	userID, err := NewUserID(11)
	require.Nil(err)

	userID2, err2 := NewUserID(112)
	require.Nil(err2)

	userID3, err3 := NewUserID(12)
	require.Nil(err3)

	require.False(userID.Equals(userID2))
	require.False(userID.Equals(userID3))

}
