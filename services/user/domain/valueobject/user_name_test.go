package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvalidUserName(t *testing.T) {
	require := require.New(t)

	userName, err := NewUserName(strings.Repeat("a", MaxUserNameLength+1))
	require.ErrorIs(err, ErrUserNameMaxLength)
	require.Len(userName, 0)

	userName, err = NewUserName(strings.Repeat("a", MinUserNameLength-1))
	require.ErrorIs(err, ErrUserNameMinLength)
	require.Len(userName, 0)

	userName, err = NewUserName(strings.Repeat("$", MaxUserNameLength))
	require.ErrorIs(err, ErrUserNameInvalidCharacter)
	require.Len(userName, 0)
}

func TestValidUserName(t *testing.T) {
	require := require.New(t)

	userName, err := NewUserName(strings.Repeat("a", MaxUserNameLength))
	require.Nil(err)
	require.NotEmpty(userName)
}

func TestEqualUserName(t *testing.T) {
	require := require.New(t)

	username, err := NewUserName(strings.Repeat("a", MaxUserNameLength))
	require.Nil(err)
	username2, err2 := NewUserName(strings.Repeat("a", MaxUserNameLength))
	require.Nil(err2)

	require.True(username.Equals(username2))

}

func TestNotEqualUserName(t *testing.T) {
	require := require.New(t)

	username, err := NewUserName(strings.Repeat("a", MaxUserNameLength))
	require.Nil(err)

	username2, err2 := NewUserName(strings.Repeat("b", MaxUserNameLength))
	require.Nil(err2)

	username3, err3 := NewUserName(strings.Repeat("a", MaxUserNameLength-1))
	require.Nil(err3)

	require.False(username.Equals(username2))
	require.False(username.Equals(username3))

}
