package uservo

import (
	"strings"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestInvalidDistrict(t *testing.T) {
	require := require.New(t)

	district, err := NewDistrict(strings.Repeat("a", MaxDistrictLength+1))
	require.ErrorIs(err, ErrDistrictMaxLength)

	require.Len(district, 0)

	district, err = NewDistrict(strings.Repeat("a", MinDistrictLength-1))
	require.ErrorIs(err, ErrDistrictMinLength)

	require.Len(district, 0)
}

func TestValidDistrict(t *testing.T) {
	require := require.New(t)

	district, err := NewDistrict(strings.Repeat("a", MaxDistrictLength))

	require.Nil(err)
	require.NotEmpty(district)
}

func TestEqualDistrict(t *testing.T) {
	require := require.New(t)

	district, err := NewDistrict(strings.Repeat("a", MaxDistrictLength))
	require.Nil(err)

	district2, err2 := NewDistrict(strings.Repeat("a", MaxDistrictLength))
	require.Nil(err2)

	require.True(district.Equals(district2))

}

func TestNotEqualDistrict(t *testing.T) {
	require := require.New(t)

	district, err := NewDistrict(strings.Repeat("a", MaxDistrictLength))
	require.Nil(err)

	district2, err2 := NewDistrict(strings.Repeat("b", MaxDistrictLength))
	require.Nil(err2)

	district3, err3 := NewDistrict(strings.Repeat("a", MaxDistrictLength-1))
	require.Nil(err3)

	require.False(district.Equals(district3))
	require.False(district.Equals(district2))

}
