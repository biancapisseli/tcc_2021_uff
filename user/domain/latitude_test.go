package userdom

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

func TestLatitudePattern(t *testing.T) {
	require := require.New(t)

	latitude, myError := NewLatitude(invalidLatitude)
	require.ErrorIs(myError, ErrLatitudeFormat)
	require.Len(latitude, 0)

	latitude2, myError := NewLatitude(invalidLatitude2)
	require.ErrorIs(myError, ErrLatitudeFormat)
	require.Len(latitude2, 0)

}

func TestLatitudeValid(t *testing.T) {
	require := require.New(t)

	latitude, myError := NewLatitude(validLatitude)
	require.Nil(myError)
	require.NotEmpty(latitude)

	latitude2, myError := NewLatitude(validLatitude2)
	require.Nil(myError)
	require.NotEmpty(latitude2)
}

func TestEqualLatitude(t *testing.T) {
	require := require.New(t)

	latitude, myError := NewLatitude(validLatitude)
	require.Nil(myError)

	latitude2, myError2 := NewLatitude(validLatitude)
	require.Nil(myError2)

	require.True(latitude.Equals(latitude2))

}

func TestNotEqualLatitude(t *testing.T) {
	require := require.New(t)

	latitude, myError := NewLatitude(validLatitude)
	require.Nil(myError)

	latitude2, myError2 := NewLatitude(validLatitude2)
	require.Nil(myError2)

	require.False(latitude.Equals(latitude2))

}
