package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	PassPattern = 9
)

func TestInvalidRawPassword(t *testing.T) {
	require := require.New(t)

	rawPassword, err := NewPasswordRaw(strings.Repeat("a", MaxRawPasswordLength+1))
	require.ErrorIs(err, ErrRawPasswordMaxLength)
	require.Len(rawPassword, 0)

	rawPassword, err = NewPasswordRaw(strings.Repeat("a", MinRawPasswordLength-1))
	require.ErrorIs(err, ErrRawPasswordMinLength)
	require.Len(rawPassword, 0)
}

func TestValidRawPassword(t *testing.T) {
	require := require.New(t)

	rawPassword, err := NewPasswordRaw("aaa#$%898")
	require.Nil(err)
	require.NotEmpty(rawPassword)
}

func TestEqualRawPassword(t *testing.T) {
	require := require.New(t)

	rawPassword, err := NewPasswordRaw(strings.Repeat("a", PassPattern))
	require.Nil(err)

	rawPassword2, err2 := NewPasswordRaw(strings.Repeat("a", PassPattern))
	require.Nil(err2)

	require.True(rawPassword.Equals(rawPassword2))
}

func TestNotEqualRawPassword(t *testing.T) {
	require := require.New(t)

	rawPassword, err := NewPasswordRaw(strings.Repeat("a", PassPattern))
	require.Nil(err)

	rawPassword2, err2 := NewPasswordRaw(strings.Repeat("b", PassPattern))
	require.Nil(err2)

	rawPassword3, err3 := NewPasswordRaw(strings.Repeat("a", PassPattern-1))
	require.Nil(err3)

	require.False(rawPassword.Equals(rawPassword3))
	require.False(rawPassword.Equals(rawPassword2))

}
