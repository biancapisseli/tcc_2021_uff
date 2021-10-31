package uservo

import (
	"strings"
	"testing"
	"github.com/stretchr/testify/require"
)

const (
	validCity = "Paraty"
)

func TestInvalidCity(t *testing.T) {
	require := require.New(t)

	city, err := NewCity(strings.Repeat("a", MaxCityLength+1))
	require.ErrorIs(err, ErrCityMaxLength)
	require.Len(city, 0)

	city, err = NewCity(strings.Repeat("a", MinCityLength-1))
	require.ErrorIs(err, ErrCityMinLength)
	require.Len(city, 0)
}

func TestValidCity(t *testing.T) {
	require := require.New(t)

	city, err := NewCity("Paraty")
	require.Nil(err)
	require.NotEmpty(city)
}

func TestEqualCity(t *testing.T) {
	require := require.New(t)

	city, err := NewCity(validCity)
	require.Nil(err)

	city2, err2 := NewCity("Paraty")
	require.Nil(err2)

	response := city.Equals(city2)
	require.True(response)

	response = city.Equals("Paraty")
	require.True(response)

}

func TestNotEqualCity(t *testing.T) {
	require := require.New(t)

	city, err := NewCity(validCity)
	require.Nil(err)

	city2, err2 := NewCity(strings.Repeat("a", MaxCityLength))
	require.Nil(err2)

	require.False(city.Equals(city2))

}
