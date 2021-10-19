package uservo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	validLongitude    = "-23.307577"
	validLongitude2   = "-3.3577"
	invalidLongitude  = "a"
	invalidLongitude2 = "20,1"
)

func TestInvalidLongitude(t *testing.T) {
	require := require.New(t)

	longitude, err := NewLongitude(invalidLongitude)
	require.ErrorIs(err, ErrLongitudeInvalidFormat)
	require.Len(longitude, 0)

	longitude2, err := NewLongitude(invalidLongitude)
	require.ErrorIs(err, ErrLongitudeInvalidFormat)
	require.Len(longitude2, 0)

}

func TestValidLongitude(t *testing.T) {
	require := require.New(t)

	longitude, err := NewLongitude(validLongitude)
	require.Nil(err)
	require.NotEmpty(longitude)

	longitude2, err := NewLongitude(validLongitude2)
	require.Nil(err)
	require.NotEmpty(longitude2)
}

func TestEqualLongitude(t *testing.T) {
	require := require.New(t)

	longitude, err := NewLongitude(validLongitude)
	require.Nil(err)

	longitude2, err2 := NewLongitude(validLongitude)
	require.Nil(err2)

	response := longitude.Equals(longitude2)
	require.True(response)

}

func TestNotEqualLongitude(t *testing.T) {
	require := require.New(t)

	longitude, err := NewLongitude(validLongitude)
	require.Nil(err)

	longitude2, err2 := NewLongitude(validLongitude2)
	require.Nil(err2)

	response := longitude.Equals(longitude2)
	require.False(response)

}
