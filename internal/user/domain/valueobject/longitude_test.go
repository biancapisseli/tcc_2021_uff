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

func TestLongitudeInvalid(t *testing.T) {
	require := require.New(t)

	longitude, myError := NewLongitude(invalidLongitude)
	require.ErrorIs(myError, ErrLongitudeInvalidFormat)
	require.Len(longitude, 0)

	longitude2, myError := NewLongitude(invalidLongitude)
	require.ErrorIs(myError, ErrLongitudeInvalidFormat)
	require.Len(longitude2, 0)

}

func TestLongitudeValid(t *testing.T) {
	require := require.New(t)

	longitude, myError := NewLongitude(validLongitude)
	require.Nil(myError)
	require.NotEmpty(longitude)

	longitude2, myError := NewLongitude(validLongitude2)
	require.Nil(myError)
	require.NotEmpty(longitude2)
}

func TestEqualLongitude(t *testing.T) {
	require := require.New(t)

	longitude, myError := NewLongitude(validLongitude)
	require.Nil(myError)

	longitude2, myError2 := NewLongitude(validLongitude)
	require.Nil(myError2)

	response := longitude.Equals(longitude2)
	require.True(response)

}

func TestNotEqualLongitude(t *testing.T) {
	require := require.New(t)

	longitude, myError := NewLongitude(validLongitude)
	require.Nil(myError)

	longitude2, myError2 := NewLongitude(validLongitude2)
	require.Nil(myError2)

	response := longitude.Equals(longitude2)
	require.False(response)

}
