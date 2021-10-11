package userdom

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidEncodedPassword(t *testing.T) {
	require := require.New(t)

	encodedPassword, myError := NewPasswordEncoded("aaa#$%898")
	require.Nil(myError)
	require.NotEmpty(encodedPassword)
}

func TestEqualEncodedPassword(t *testing.T) {
	require := require.New(t)

	encodedPassword, myError := NewPasswordEncoded(strings.Repeat("a", PassPattern))
	require.Nil(myError)

	encodedPassword2, myError2 := NewPasswordEncoded(strings.Repeat("a", PassPattern))
	require.Nil(myError2)

	require.True(encodedPassword.Equals(encodedPassword2))
}

func TestNotEqualEncodedPassword(t *testing.T) {
	require := require.New(t)

	encodedPassword, myError := NewPasswordEncoded(strings.Repeat("a", PassPattern))
	require.Nil(myError)

	encodedPassword2, myError2 := NewPasswordEncoded(strings.Repeat("b", PassPattern))
	require.Nil(myError2)

	encodedPassword3, myError3 := NewPasswordEncoded(strings.Repeat("a", PassPattern-1))
	require.Nil(myError3)

	require.False(encodedPassword.Equals(encodedPassword3))
	require.False(encodedPassword.Equals(encodedPassword2))

}
