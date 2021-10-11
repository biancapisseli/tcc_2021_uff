package userdom

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	PassPattern = 9
)

func TestRawPasswordMaxLength(t *testing.T) {
	require := require.New(t)

	rawPassword, myError := NewPasswordRaw(strings.Repeat("a", maxRawPasswordLength+1))
	require.ErrorIs(myError, ErrRawPasswordMaxLength)
	require.Len(rawPassword, 0)
}

func TestRawPasswordMinLength(t *testing.T) {
	require := require.New(t)

	rawPassword, myError := NewPasswordRaw(strings.Repeat("a", minRawPasswordLength-1))
	require.ErrorIs(myError, ErrRawPasswordMinLength)
	require.Len(rawPassword, 0)
}

func TestValidRawPassword(t *testing.T) {
	require := require.New(t)

	rawPassword, myError := NewPasswordRaw("aaa#$%898")
	require.Nil(myError)
	require.NotEmpty(rawPassword)
}

func TestEqualRawPassword(t *testing.T) {
	require := require.New(t)

	rawPassword, myError := NewPasswordRaw(strings.Repeat("a", PassPattern))
	require.Nil(myError)

	rawPassword2, myError2 := NewPasswordRaw(strings.Repeat("a", PassPattern))
	require.Nil(myError2)

	require.True(rawPassword.Equals(rawPassword2))
}

func TestNotEqualRawPassword(t *testing.T) {
	require := require.New(t)

	rawPassword, myError := NewPasswordRaw(strings.Repeat("a", PassPattern))
	require.Nil(myError)

	rawPassword2, myError2 := NewPasswordRaw(strings.Repeat("b", PassPattern))
	require.Nil(myError2)

	rawPassword3, myError3 := NewPasswordRaw(strings.Repeat("a", PassPattern-1))
	require.Nil(myError3)

	require.False(rawPassword.Equals(rawPassword3))
	require.False(rawPassword.Equals(rawPassword2))

}
