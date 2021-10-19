package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidEncodedPassword(t *testing.T) {
	require := require.New(t)

	encodedPassword, err := NewPasswordEncoded("aaa#$%898")
	require.Nil(err)
	require.NotEmpty(encodedPassword)
}

func TestEqualEncodedPassword(t *testing.T) {
	require := require.New(t)

	encodedPassword, err := NewPasswordEncoded(strings.Repeat("a", PassPattern))
	require.Nil(err)

	encodedPassword2, err2 := NewPasswordEncoded(strings.Repeat("a", PassPattern))
	require.Nil(err2)

	require.True(encodedPassword.Equals(encodedPassword2))
}

func TestNotEqualEncodedPassword(t *testing.T) {
	require := require.New(t)

	encodedPassword, err := NewPasswordEncoded(strings.Repeat("a", PassPattern))
	require.Nil(err)

	encodedPassword2, err2 := NewPasswordEncoded(strings.Repeat("b", PassPattern))
	require.Nil(err2)

	encodedPassword3, err3 := NewPasswordEncoded(strings.Repeat("a", PassPattern-1))
	require.Nil(err3)

	require.False(encodedPassword.Equals(encodedPassword3))
	require.False(encodedPassword.Equals(encodedPassword2))

}
