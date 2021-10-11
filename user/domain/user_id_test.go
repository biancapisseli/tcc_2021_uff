package userdom

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserIDValid(t *testing.T) {
	require := require.New(t)

	id, myError := NewUserID(5)
	require.Nil(myError)
	require.NotEmpty(id)
}

func TestUserIDInvalid(t *testing.T) {
	require := require.New(t)

	id, myError := NewUserID(-1)
	require.ErrorIs(myError, ErrInvalidUserID)
	require.ElementsMatch(id, 0)
}

func TestEqualUserID(t *testing.T) {
	require := require.New(t)

	userID, myError := NewUserID(11)
	require.Nil(myError)

	userID2, myError2 := NewUserID(11)
	require.Nil(myError2)

	require.True(userID.Equals(userID2))
}

func TestNotEqualUserID(t *testing.T) {
	require := require.New(t)

	userID, myError := NewUserID(11)
	require.Nil(myError)

	userID2, myError2 := NewUserID(112)
	require.Nil(myError2)

	userID3, myError3 := NewUserID(12)
	require.Nil(myError3)

	require.False(userID.Equals(userID2))
	require.False(userID.Equals(userID3))

}
