package uservo

import (
	"strings"
	"testing"
	"github.com/stretchr/testify/require"
)

const (
	EMAIL_HOSTMAME = "@gmail.com"
)

func TestInvalidEmail(t *testing.T) {
	require := require.New(t)

	email, err := NewEmail((strings.Repeat("a", MaxEmailLength) + EMAIL_HOSTMAME))
	require.ErrorIs(err, ErrEmailMaxLength)
	require.Len(email, 0)

	email, err = NewEmail((strings.Repeat("@", MaxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME))
	require.ErrorIs(err, ErrEmailInvalidFormat)
	require.Len(email, 0)

	email, err = NewEmail((strings.Repeat("a", MaxEmailLength-len(EMAIL_HOSTMAME)) + "@@fdsffds"))
	require.ErrorIs(err, ErrEmailInvalidFormat)
	require.Len(email, 0)

	email, err = NewEmail((strings.Repeat("a", MaxEmailLength-len(EMAIL_HOSTMAME))))
	require.ErrorIs(err, ErrEmailInvalidFormat)
	require.Len(email, 0)

}

func TestValidEmail(t *testing.T) {
	require := require.New(t)

	email, err := NewEmail((strings.Repeat("a", MaxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME))
	require.Nil(err)

	require.NotEmpty(email)

}

func TestEqualEmail(t *testing.T) {
	require := require.New(t)

	email, err := NewEmail((strings.Repeat("a", MaxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME))
	require.Nil(err)

	email2, err2 := NewEmail((strings.Repeat("a", MaxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME))
	require.Nil(err2)

	require.True(email.Equals(email2))

}

func TestNotEqualEmail(t *testing.T) {
	require := require.New(t)

	email, err := NewEmail((strings.Repeat("a", MaxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME))
	require.Nil(err)

	email2, err2 := NewEmail((strings.Repeat("b", MaxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME))
	require.Nil(err2)

	email3, err3 := NewEmail((strings.Repeat("a", MaxEmailLength-len(EMAIL_HOSTMAME)-1) + EMAIL_HOSTMAME))
	require.Nil(err3)

	require.False(email.Equals(email2))
	require.False(email.Equals(email3))

}
