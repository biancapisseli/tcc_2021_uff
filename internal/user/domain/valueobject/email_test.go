package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	EMAIL_HOSTMAME = "@gmail.com"
)

func TestEmailMaxLength(t *testing.T) {
	require := require.New(t)

	email, myError := NewEmail((strings.Repeat("a", MaxEmailLength) + EMAIL_HOSTMAME))
	require.ErrorIs(myError, ErrEmailMaxLength)
	require.Len(email, 0)

}

func TestEmailPattern(t *testing.T) {
	require := require.New(t)

	email, myError := NewEmail((strings.Repeat("@", MaxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME))
	require.ErrorIs(myError, ErrEmailInvalidFormat)
	require.Len(email, 0)

	email, myError = NewEmail((strings.Repeat("a", MaxEmailLength-len(EMAIL_HOSTMAME)) + "@@fdsffds"))
	require.ErrorIs(myError, ErrEmailInvalidFormat)
	require.Len(email, 0)

	email, myError = NewEmail((strings.Repeat("a", MaxEmailLength-len(EMAIL_HOSTMAME))))
	require.ErrorIs(myError, ErrEmailInvalidFormat)
	require.Len(email, 0)
}

func TestValidEmail(t *testing.T) {
	require := require.New(t)

	email, myError := NewEmail((strings.Repeat("a", MaxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME))
	require.Nil(myError)

	require.NotEmpty(email)

}

func TestEqualEmail(t *testing.T) {
	require := require.New(t)

	email, myError := NewEmail((strings.Repeat("a", MaxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME))
	require.Nil(myError)

	email2, myError2 := NewEmail((strings.Repeat("a", MaxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME))
	require.Nil(myError2)

	require.True(email.Equals(email2))

}

func TestNotEqualEmail(t *testing.T) {
	require := require.New(t)

	email, myError := NewEmail((strings.Repeat("a", MaxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME))
	require.Nil(myError)

	email2, myError2 := NewEmail((strings.Repeat("b", MaxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME))
	require.Nil(myError2)

	email3, myError3 := NewEmail((strings.Repeat("a", MaxEmailLength-len(EMAIL_HOSTMAME)-1) + EMAIL_HOSTMAME))
	require.Nil(myError3)

	require.False(email.Equals(email2))
	require.False(email.Equals(email3))

}
