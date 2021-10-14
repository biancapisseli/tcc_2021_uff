package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	validCity = "Paraty"
)

func TestCityMaxLength(t *testing.T) {
	require := require.New(t)

	city, myError := NewCity(strings.Repeat("a", maxCityLength+1))
	require.ErrorIs(myError, ErrCityMaxLength)
	require.Len(city, 0)
}

func TestValidCity(t *testing.T) {
	require := require.New(t)

	city, myError := NewCity(validCity)
	require.Nil(myError)
	require.NotEmpty(city)
}

func TestEqualCity(t *testing.T) {
	require := require.New(t)

	city, myError := NewCity(validCity)
	require.Nil(myError)

	city2, myError2 := NewCity(validCity)
	require.Nil(myError2)

	response := city.Equals(city2)
	require.True(response)

}

func TestNotEqualCity(t *testing.T) {
	require := require.New(t)

	city, myError := NewCity(validCity)
	require.Nil(myError)

	city2, myError2 := NewCity(strings.Repeat("a", maxCityLength))
	require.Nil(myError2)

	require.False(city.Equals(city2))

}
