package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidEncodedPassword(t *testing.T) {
	require := require.New(t)

	encodedPassword := NewPasswordEncoded("aaa#$%898")
	require.Equal("aaa#$%898", encodedPassword.String())
}

func TestEqualEncodedPassword(t *testing.T) {
	require := require.New(t)

	encodedPassword := NewPasswordEncoded(strings.Repeat("a", PassPattern))
	encodedPassword2 := NewPasswordEncoded(strings.Repeat("a", PassPattern))

	require.True(encodedPassword.Equals(encodedPassword2))
}

func TestNotEqualEncodedPassword(t *testing.T) {
	require := require.New(t)

	encodedPassword := NewPasswordEncoded(strings.Repeat("a", PassPattern))
	encodedPassword2 := NewPasswordEncoded(strings.Repeat("b", PassPattern))
	encodedPassword3 := NewPasswordEncoded(strings.Repeat("a", PassPattern-1))

	require.False(encodedPassword.Equals(encodedPassword3))
	require.False(encodedPassword.Equals(encodedPassword2))

}
