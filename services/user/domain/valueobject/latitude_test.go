package uservo

import (
	"testing"
	"github.com/stretchr/testify/require"
)

const (
	validLatitude    = "-23.307577"
	validLatitude2   = "-3.3577"
	invalidLatitude  = "a"
	invalidLatitude2 = "20,1"
)

func TestInvalidLatitude(t *testing.T) {

	require := require.New(t)

	latitude, err := NewLatitude(invalidLatitude)
	require.ErrorIs(err, ErrLatitudeInvalidFormat)
	require.Len(latitude, 0)

	latitude2, err := NewLatitude(invalidLatitude2)
	require.ErrorIs(err, ErrLatitudeInvalidFormat)
	require.Len(latitude2, 0)

}

func TestValidLatitude(t *testing.T) {

	require := require.New(t)

	latitude, err := NewLatitude(validLatitude)
	require.Nil(err)
	require.NotEmpty(latitude)

	latitude2, err := NewLatitude(validLatitude2)
	require.Nil(err)
	require.NotEmpty(latitude2)
}

func TestEqualLatitude(t *testing.T) {
	require := require.New(t)

	latitude, err := NewLatitude(validLatitude)
	require.Nil(err)

	latitude2, err2 := NewLatitude(validLatitude)
	require.Nil(err2)

	require.True(latitude.Equals(latitude2))

}

func TestNotEqualLatitude(t *testing.T) {
	require := require.New(t)

	latitude, err := NewLatitude(validLatitude)
	require.Nil(err)

	latitude2, err2 := NewLatitude(validLatitude2)
	require.Nil(err2)

	require.False(latitude.Equals(latitude2))

}
