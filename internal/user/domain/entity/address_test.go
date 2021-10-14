package userent

import (
	"encoding/json"
	"strings"
	"testing"

	uservo "ifoodish-store/internal/domain/valueobject"

	"github.com/stretchr/testify/require"
)

var (
	validAddress = Address{
		Street:     "Street ABCD",
		District:   "Espirito Santo",
		City:       "Jose dos Campos",
		State:      "Rio de Janeiro",
		Complement: "Complement",
		Number:     "11111",
		Zipcode:    "23970000",
		Latitude:   "-23.307577",
		Longitude:  "-44.754146",
	}
	invalidAddress1 = Address{
		Street:     uservo.Street(strings.Repeat("a", minStreetLength-1)),
		District:   uservo.District(strings.Repeat("a", minDistrictLength-1)),
		City:       uservo.City(strings.Repeat("a", maxCityLength+1)),
		State:      uservo.State((strings.Repeat("a", maxStateLength+1))),
		Complement: uservo.Complement((strings.Repeat("a", maxComplementLength+1))),
		Number:     uservo.Number((strings.Repeat("1", maxNumberLength+1))),
		Zipcode:    uservo.Zipcode((strings.Repeat("1", zipcodeLength+1))),
		Latitude:   uservo.Latitude((strings.Repeat("1", 5))),
		Longitude:  uservo.Longitude((strings.Repeat("1", 5))),
	}
	invalidAddress2 = Address{
		Street:     "Street ABCD",
		District:   uservo.District(strings.Repeat("a", minDistrictLength-1)),
		City:       uservo.City(strings.Repeat("a", maxCityLength+1)),
		State:      uservo.State((strings.Repeat("a", maxStateLength+1))),
		Complement: uservo.Complement((strings.Repeat("a", maxComplementLength+1))),
		Number:     uservo.Number((strings.Repeat("1", maxNumberLength+1))),
		Zipcode:    uservo.Zipcode((strings.Repeat("1", zipcodeLength+1))),
		Latitude:   uservo.Latitude((strings.Repeat("1", 5))),
		Longitude:  uservo.Longitude((strings.Repeat("1", 5))),
	}
	invalidAddress3 = Address{
		Street:     "Street ABCD",
		District:   "District",
		City:       uservo.City(strings.Repeat("a", maxCityLength+1)),
		State:      uservo.State((strings.Repeat("a", maxStateLength+1))),
		Complement: uservo.Complement((strings.Repeat("a", maxComplementLength+1))),
		Number:     uservo.Number((strings.Repeat("1", maxNumberLength+1))),
		Zipcode:    uservo.Zipcode((strings.Repeat("1", zipcodeLength+1))),
		Latitude:   uservo.Latitude((strings.Repeat("1", 5))),
		Longitude:  uservo.Longitude((strings.Repeat("1", 5))),
	}
	invalidAddress4 = Address{
		Street:     "Street ABCD",
		District:   "District",
		City:       "City",
		State:      uservo.State((strings.Repeat("a", maxStateLength+1))),
		Complement: uservo.Complement((strings.Repeat("a", maxComplementLength+1))),
		Number:     uservo.Number((strings.Repeat("1", maxNumberLength+1))),
		Zipcode:    uservo.Zipcode((strings.Repeat("1", zipcodeLength+1))),
		Latitude:   uservo.Latitude((strings.Repeat("1", 5))),
		Longitude:  uservo.Longitude((strings.Repeat("1", 5))),
	}
	invalidAddress5 = Address{
		Street:     "Street ABCD",
		District:   "District",
		City:       "City",
		State:      "State",
		Complement: uservo.Complement((strings.Repeat("a", maxComplementLength+1))),
		Number:     uservo.Number((strings.Repeat("1", maxNumberLength+1))),
		Zipcode:    uservo.Zipcode((strings.Repeat("1", zipcodeLength+1))),
		Latitude:   uservo.Latitude((strings.Repeat("1", 5))),
		Longitude:  uservo.Longitude((strings.Repeat("1", 5))),
	}
	invalidAddress6 = Address{
		Street:     "Street ABCD",
		District:   "District",
		City:       "City",
		State:      "State",
		Complement: "Complement",
		Number:     uservo.Number((strings.Repeat("1", maxNumberLength+1))),
		Zipcode:    uservo.Zipcode((strings.Repeat("1", zipcodeLength+1))),
		Latitude:   uservo.Latitude((strings.Repeat("1", 5))),
		Longitude:  uservo.Longitude((strings.Repeat("1", 5))),
	}
	invalidAddress7 = Address{
		Street:     "Street ABCD",
		District:   "District",
		City:       "City",
		State:      "State",
		Complement: "Complement",
		Number:     "11111",
		Zipcode:    uservo.Zipcode((strings.Repeat("1", zipcodeLength+1))),
		Latitude:   uservo.Latitude((strings.Repeat("1", 5))),
		Longitude:  uservo.Longitude((strings.Repeat("1", 5))),
	}
	invalidAddress8 = Address{
		Street:     "Street ABCD",
		District:   "District",
		City:       "City",
		State:      "State",
		Complement: "Complement",
		Number:     "11111",
		Zipcode:    "23970000",
		Latitude:   uservo.Latitude((strings.Repeat("1", 5))),
		Longitude:  uservo.Longitude((strings.Repeat("1", 5))),
	}
	invalidAddress9 = Address{
		Street:     "Street ABCD",
		District:   "District",
		City:       "City",
		State:      "State",
		Complement: "Complement",
		Number:     "11111",
		Zipcode:    "23970000",
		Latitude:   "-23.307577",
		Longitude:  uservo.Longitude((strings.Repeat("1", 5))),
	}
	validRegisteredAddress = RegisteredAddress{
		Address: validAddress,
		ID:      50,
	}
	invalidRegisteredAddress1 = RegisteredAddress{
		Address: invalidAddress1,
		ID:      -1,
	}
	invalidRegisteredAddress2 = RegisteredAddress{
		Address: validAddress,
		ID:      -1,
	}
)

func TestAddressValid(t *testing.T) {
	require := require.New(t)
	address, myError := NewAddress(validAddress)
	require.Nil(myError)
	require.NotEmpty(address)
}

func TestAddressInvalid(t *testing.T) {
	require := require.New(t)

	address, myError := NewAddress(invalidAddress1)
	require.NotNil(myError)
	require.Nil(address)

	address, myError = NewAddress(invalidAddress2)
	require.NotNil(myError)
	require.Nil(address)

	address, myError = NewAddress(invalidAddress3)
	require.NotNil(myError)
	require.Nil(address)

	address, myError = NewAddress(invalidAddress4)
	require.NotNil(myError)
	require.Nil(address)

	address, myError = NewAddress(invalidAddress5)
	require.NotNil(myError)
	require.Nil(address)

	address, myError = NewAddress(invalidAddress6)
	require.NotNil(myError)
	require.Nil(address)

	address, myError = NewAddress(invalidAddress7)
	require.NotNil(myError)
	require.Nil(address)

	address, myError = NewAddress(invalidAddress8)
	require.NotNil(myError)
	require.Nil(address)

	address, myError = NewAddress(invalidAddress9)
	require.NotNil(myError)
	require.Nil(address)
}

func TestRegisteredAddressValid(t *testing.T) {
	require := require.New(t)
	address, myError := NewRegisteredAddress(validRegisteredAddress)
	require.Nil(myError)
	require.NotEmpty(address)
}

func TestRegisteredAddressInvalid(t *testing.T) {
	require := require.New(t)

	address, myError := NewRegisteredAddress(invalidRegisteredAddress1)
	require.NotNil(myError)
	require.Nil(address)

	address, myError = NewRegisteredAddress(invalidRegisteredAddress2)
	require.NotNil(myError)
	require.Nil(address)

}

func TestJSONUnmarshallingAddressSuccess(t *testing.T) {
	require := require.New(t)

	var address *Address
	err := json.Unmarshal([]byte(`
	{
		"Street":     "Street ABCD",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.Nil(err)
	require.True(address.Street.Equals("Street ABCD"))
	require.True(address.District.Equals("District"))
	require.True(address.City.Equals("City"))
	require.True(address.State.Equals("State"))
	require.True(address.Complement.Equals("Complement"))
	require.True(address.Number.Equals("11111"))
	require.True(address.Zipcode.Equals("23970000"))
	require.True(address.Latitude.Equals("-23.307577"))
	require.True(address.Longitude.Equals("-44.754146"))
}

func TestJSONUnmarshallingAddressFail(t *testing.T) {
	require := require.New(t)
	var address *Address

	// forçando teste do unmarshal
	err := address.UnmarshalJSON([]byte(`
		{
			"Street":     "Street ABCD",
			"District":   "District",
			"City":       "City",
			"State":      "State",
			"Complement": "Complement",
			"Number":     "11111",
			"Zipcode":    "23970000",
			"Latitude":   "-23.307577",
			"Longitude":  "-44.754146
		}
	`))
	require.NotNil(err)

	address = nil
	err = json.Unmarshal([]byte(`
	{
		"Street":     "",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, ErrStreetMinLength)

}
