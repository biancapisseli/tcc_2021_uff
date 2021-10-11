package userdom

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPhoneMaxLength(t *testing.T) {
	require := require.New(t)

	Phone, myError := NewPhone(strings.Repeat("1", maxPhoneLength+1))
	require.ErrorIs(myError, ErrPhoneMaxLength)
	require.Len(Phone, 0)
}

func TestPhoneMinLength(t *testing.T) {
	require := require.New(t)

	Phone, myError := NewPhone(strings.Repeat("1", minPhoneLength-1))
	require.ErrorIs(myError, ErrPhoneMinLength)
	require.Len(Phone, 0)
}

func TestValidPhone(t *testing.T) {
	require := require.New(t)

	Phone, myError := NewPhone("24999224073")
	require.Nil(myError)
	require.NotEmpty(Phone)
}

func TestEqualPhone(t *testing.T) {
	require := require.New(t)

	Phone, myError := NewPhone("24999224073")
	require.Nil(myError)

	Phone2, myError2 := NewPhone("24999224073")
	require.Nil(myError2)

	require.True(Phone.Equals(Phone2))
}

func TestNotEqualPhone(t *testing.T) {
	require := require.New(t)

	Phone, myError := NewPhone("24999224073")
	require.Nil(myError)

	Phone2, myError2 := NewPhone("24999224083")
	require.Nil(myError2)

	Phone3, myError3 := NewPhone("2499224073")
	require.Nil(myError3)

	require.False(Phone.Equals(Phone3))
	require.False(Phone.Equals(Phone2))

}
