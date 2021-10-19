package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvalidPhone(t *testing.T) {
	require := require.New(t)

	Phone, err := NewPhone(strings.Repeat("1", MaxPhoneLength+1))
	require.ErrorIs(err, ErrPhoneMaxLength)
	require.Len(Phone, 0)

	Phone, err = NewPhone(strings.Repeat("1", MinPhoneLength-1))
	require.ErrorIs(err, ErrPhoneMinLength)
	require.Len(Phone, 0)
}

func TestValidPhone(t *testing.T) {
	require := require.New(t)

	Phone, err := NewPhone("24999224073")
	require.Nil(err)
	require.NotEmpty(Phone)
}

func TestEqualPhone(t *testing.T) {
	require := require.New(t)

	Phone, err := NewPhone("24999224073")
	require.Nil(err)

	Phone2, err2 := NewPhone("24999224073")
	require.Nil(err2)

	require.True(Phone.Equals(Phone2))
}

func TestNotEqualPhone(t *testing.T) {
	require := require.New(t)

	Phone, err := NewPhone("24999224073")
	require.Nil(err)

	Phone2, err2 := NewPhone("24999224083")
	require.Nil(err2)

	Phone3, err3 := NewPhone("2499224073")
	require.Nil(err3)

	require.False(Phone.Equals(Phone3))
	require.False(Phone.Equals(Phone2))

}
