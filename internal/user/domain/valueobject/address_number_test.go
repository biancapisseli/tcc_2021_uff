package uservo

import (
	"strings"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestInvalidAddressNumber(t *testing.T) {
	require := require.New(t)

	AddressNumber, err := NewAddressNumber(strings.Repeat("1", MaxAddressNumberLength+1))
	require.ErrorIs(err, ErrAddressNumberMaxLength)
	require.Len(AddressNumber, 0)

	AddressNumber, err = NewAddressNumber(strings.Repeat("1", MinAddressNumberLength-1))
	require.ErrorIs(err, ErrAddressNumberMinLength)
	require.Len(AddressNumber, 0)
}

func TestValidAddressNumber(t *testing.T) {
	require := require.New(t)

	AddressNumber, err := NewAddressNumber(strings.Repeat("1", MaxAddressNumberLength-1))
	require.Nil(err)
	require.NotEmpty(AddressNumber)
}

func TestEqualAddressNumber(t *testing.T) {
	require := require.New(t)

	AddressNumber, err := NewAddressNumber(strings.Repeat("1", MaxAddressNumberLength-1))
	require.Nil(err)

	AddressNumber2, err2 := NewAddressNumber(strings.Repeat("1", MaxAddressNumberLength-1))
	require.Nil(err2)

	require.True(AddressNumber.Equals(AddressNumber2))

}

func TestNotEqualAddressNumber(t *testing.T) {
	require := require.New(t)

	AddressNumber, err := NewAddressNumber(strings.Repeat("1", MaxAddressNumberLength-1))
	require.Nil(err)

	AddressNumber2, err2 := NewAddressNumber(strings.Repeat("2", MaxAddressNumberLength-1))
	require.Nil(err2)

	AddressNumber3, err3 := NewAddressNumber(strings.Repeat("1", MaxAddressNumberLength))
	require.Nil(err3)

	require.False(AddressNumber.Equals(AddressNumber3))
	require.False(AddressNumber.Equals(AddressNumber2))

}
