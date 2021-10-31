package userent_test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"

	"github.com/carlmjohnson/resperr"

	"github.com/stretchr/testify/require"
)

type addressTestCase struct {
	street     string
	district   string
	city       string
	state      string
	complement string
	number     string
	zipcode    string
	latitude   string
	longitude  string

	expectedErr error
}

type registeredAddressTestCase struct {
	addressTestCase
	id int64
}

var (
	validAddressTestCase = addressTestCase{
		street:     "Street ABCD",
		district:   "Espirito Santo",
		city:       "Jose dos Campos",
		state:      "Rio de Janeiro",
		complement: "Complement",
		number:     "11111",
		zipcode:    "23970000",
		latitude:   "-23.307577",
		longitude:  "-44.754146",

		expectedErr: nil,
	}
)

func addressTestCaseCompare(
	require *require.Assertions,
	address userent.Address,
	tc addressTestCase,
) {
	require.Equal(tc.street, address.Street.String())
	require.Equal(tc.district, address.District.String())
	require.Equal(tc.city, address.City.String())
	require.Equal(tc.state, address.State.String())
	require.Equal(tc.complement, address.Complement.String())
	require.Equal(tc.number, address.Number.String())
	require.Equal(tc.zipcode, address.Zipcode.String())
	require.Equal(tc.latitude, address.Latitude.String())
	require.Equal(tc.longitude, address.Longitude.String())
}

func TestAddressValid(t *testing.T) {
	require := require.New(t)
	address, err := userent.NewAddress(
		validAddressTestCase.street,
		validAddressTestCase.district,
		validAddressTestCase.city,
		validAddressTestCase.state,
		validAddressTestCase.complement,
		validAddressTestCase.number,
		validAddressTestCase.zipcode,
		validAddressTestCase.latitude,
		validAddressTestCase.longitude,
	)
	require.Nil(err)
	addressTestCaseCompare(require, address, validAddressTestCase)
}

func TestAddressInvalid(t *testing.T) {
	require := require.New(t)

	addresses := []addressTestCase{{
		street:      strings.Repeat("a", uservo.MinStreetLength-1),
		district:    validAddressTestCase.district,
		city:        validAddressTestCase.city,
		state:       validAddressTestCase.state,
		complement:  validAddressTestCase.complement,
		number:      validAddressTestCase.number,
		zipcode:     validAddressTestCase.zipcode,
		latitude:    validAddressTestCase.latitude,
		longitude:   validAddressTestCase.longitude,
		expectedErr: uservo.ErrStreetMinLength,
	}, {
		street:      strings.Repeat("a", uservo.MaxStreetLength+1),
		district:    validAddressTestCase.district,
		city:        validAddressTestCase.city,
		state:       validAddressTestCase.state,
		complement:  validAddressTestCase.complement,
		number:      validAddressTestCase.number,
		zipcode:     validAddressTestCase.zipcode,
		latitude:    validAddressTestCase.latitude,
		longitude:   validAddressTestCase.longitude,
		expectedErr: uservo.ErrStreetMaxLength,
	}, {
		street:      validAddressTestCase.street,
		district:    strings.Repeat("a", uservo.MinDistrictLength-1),
		city:        validAddressTestCase.city,
		state:       validAddressTestCase.state,
		complement:  validAddressTestCase.complement,
		number:      validAddressTestCase.number,
		zipcode:     validAddressTestCase.zipcode,
		latitude:    validAddressTestCase.latitude,
		longitude:   validAddressTestCase.longitude,
		expectedErr: uservo.ErrDistrictMinLength,
	}, {
		street:      validAddressTestCase.street,
		district:    strings.Repeat("a", uservo.MaxDistrictLength+1),
		city:        validAddressTestCase.city,
		state:       validAddressTestCase.state,
		complement:  validAddressTestCase.complement,
		number:      validAddressTestCase.number,
		zipcode:     validAddressTestCase.zipcode,
		latitude:    validAddressTestCase.latitude,
		longitude:   validAddressTestCase.longitude,
		expectedErr: uservo.ErrDistrictMaxLength,
	}, {
		street:      validAddressTestCase.street,
		district:    validAddressTestCase.district,
		city:        strings.Repeat("a", uservo.MinCityLength-1),
		state:       validAddressTestCase.state,
		complement:  validAddressTestCase.complement,
		number:      validAddressTestCase.number,
		zipcode:     validAddressTestCase.zipcode,
		latitude:    validAddressTestCase.latitude,
		longitude:   validAddressTestCase.longitude,
		expectedErr: uservo.ErrCityMinLength,
	}, {
		street:      validAddressTestCase.street,
		district:    validAddressTestCase.district,
		city:        strings.Repeat("a", uservo.MaxCityLength+1),
		state:       validAddressTestCase.state,
		complement:  validAddressTestCase.complement,
		number:      validAddressTestCase.number,
		zipcode:     validAddressTestCase.zipcode,
		latitude:    validAddressTestCase.latitude,
		longitude:   validAddressTestCase.longitude,
		expectedErr: uservo.ErrCityMaxLength,
	}, {
		street:      validAddressTestCase.street,
		district:    validAddressTestCase.district,
		city:        validAddressTestCase.city,
		state:       strings.Repeat("a", uservo.MinStateLength-1),
		complement:  validAddressTestCase.complement,
		number:      validAddressTestCase.number,
		zipcode:     validAddressTestCase.zipcode,
		latitude:    validAddressTestCase.latitude,
		longitude:   validAddressTestCase.longitude,
		expectedErr: uservo.ErrStateMinLength,
	}, {
		street:      validAddressTestCase.street,
		district:    validAddressTestCase.district,
		city:        validAddressTestCase.city,
		state:       strings.Repeat("a", uservo.MaxStateLength+1),
		complement:  validAddressTestCase.complement,
		number:      validAddressTestCase.number,
		zipcode:     validAddressTestCase.zipcode,
		latitude:    validAddressTestCase.latitude,
		longitude:   validAddressTestCase.longitude,
		expectedErr: uservo.ErrStateMaxLength,
	}, {
		street:      validAddressTestCase.street,
		district:    validAddressTestCase.district,
		city:        validAddressTestCase.city,
		state:       validAddressTestCase.state,
		complement:  strings.Repeat("a", uservo.MaxComplementLength+1),
		number:      validAddressTestCase.number,
		zipcode:     validAddressTestCase.zipcode,
		latitude:    validAddressTestCase.latitude,
		longitude:   validAddressTestCase.longitude,
		expectedErr: uservo.ErrComplementMaxLength,
	}, {
		street:      validAddressTestCase.street,
		district:    validAddressTestCase.district,
		city:        validAddressTestCase.city,
		state:       validAddressTestCase.state,
		complement:  validAddressTestCase.complement,
		number:      strings.Repeat("a", uservo.MinAddressNumberLength-1),
		zipcode:     validAddressTestCase.zipcode,
		latitude:    validAddressTestCase.latitude,
		longitude:   validAddressTestCase.longitude,
		expectedErr: uservo.ErrAddressNumberMinLength,
	}, {
		street:      validAddressTestCase.street,
		district:    validAddressTestCase.district,
		city:        validAddressTestCase.city,
		state:       validAddressTestCase.state,
		complement:  validAddressTestCase.complement,
		number:      strings.Repeat("a", uservo.MaxAddressNumberLength+1),
		zipcode:     validAddressTestCase.zipcode,
		latitude:    validAddressTestCase.latitude,
		longitude:   validAddressTestCase.longitude,
		expectedErr: uservo.ErrAddressNumberMaxLength,
	}, {
		street:      validAddressTestCase.street,
		district:    validAddressTestCase.district,
		city:        validAddressTestCase.city,
		state:       validAddressTestCase.state,
		complement:  validAddressTestCase.complement,
		number:      validAddressTestCase.number,
		zipcode:     strings.Repeat("1", uservo.ZipcodeLength+1),
		latitude:    validAddressTestCase.latitude,
		longitude:   validAddressTestCase.longitude,
		expectedErr: uservo.ErrZipcodeLength,
	}, {
		street:      validAddressTestCase.street,
		district:    validAddressTestCase.district,
		city:        validAddressTestCase.city,
		state:       validAddressTestCase.state,
		complement:  validAddressTestCase.complement,
		number:      validAddressTestCase.number,
		zipcode:     strings.Repeat("a", uservo.ZipcodeLength),
		latitude:    validAddressTestCase.latitude,
		longitude:   validAddressTestCase.longitude,
		expectedErr: uservo.ErrZipcodeNotNumeric,
	}, {
		street:      validAddressTestCase.street,
		district:    validAddressTestCase.district,
		city:        validAddressTestCase.city,
		state:       validAddressTestCase.state,
		complement:  validAddressTestCase.complement,
		number:      validAddressTestCase.number,
		zipcode:     validAddressTestCase.zipcode,
		latitude:    strings.Repeat("a", 5),
		longitude:   validAddressTestCase.longitude,
		expectedErr: uservo.ErrLatitudeInvalidFormat,
	}, {
		street:      validAddressTestCase.street,
		district:    validAddressTestCase.district,
		city:        validAddressTestCase.city,
		state:       validAddressTestCase.state,
		complement:  validAddressTestCase.complement,
		number:      validAddressTestCase.number,
		zipcode:     validAddressTestCase.zipcode,
		latitude:    validAddressTestCase.latitude,
		longitude:   strings.Repeat("a", 5),
		expectedErr: uservo.ErrLongitudeInvalidFormat,
	}}

	for i, it := range addresses {
		_, err := userent.NewAddress(
			it.street, it.district, it.city, it.state, it.complement,
			it.number, it.zipcode, it.latitude, it.longitude,
		)
		require.ErrorIs(err, it.expectedErr, "index %d", i)
	}

}

func TestRegisteredAddressValid(t *testing.T) {
	require := require.New(t)

	address, err := userent.NewAddress(
		validAddressTestCase.street,
		validAddressTestCase.district,
		validAddressTestCase.city,
		validAddressTestCase.state,
		validAddressTestCase.complement,
		validAddressTestCase.number,
		validAddressTestCase.zipcode,
		validAddressTestCase.latitude,
		validAddressTestCase.longitude,
	)
	require.Nil(err)
	addressTestCaseCompare(require, address, validAddressTestCase)

	regAddress, err := userent.NewRegisteredAddress(50, address)
	require.Nil(err)
	addressTestCaseCompare(require, regAddress.Address, validAddressTestCase)
	require.Equal("50", regAddress.ID.String())
}

func TestRegisteredAddressInvalidID(t *testing.T) {
	require := require.New(t)

	for _, it := range []registeredAddressTestCase{{
		addressTestCase: validAddressTestCase,
		id:              0,
	}, {
		addressTestCase: validAddressTestCase,
		id:              -10,
	}} {
		address, err := userent.NewAddress(
			it.street, it.district, it.city, it.state, it.complement,
			it.number, it.zipcode, it.latitude, it.longitude,
		)
		require.Nil(err)
		addressTestCaseCompare(require, address, validAddressTestCase)

		_, err = userent.NewRegisteredAddress(it.id, address)
		require.ErrorIs(err, uservo.ErrInvalidAddressID)
	}
}

func TestJSONUnmarshallingAddressSuccess(t *testing.T) {
	require := require.New(t)

	var address *userent.Address
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
	var address *userent.Address

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
	require.ErrorIs(err, uservo.ErrStreetMinLength)

}

//////////////////////////

func TestJSONUnmarshallingRegisteredAddressSuccess(t *testing.T) {
	require := require.New(t)

	var address userent.RegisteredAddress
	err := address.UnmarshalJSON([]byte(`
	{
		"id":         50,	
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
	`))
	require.Nil(err)
	require.True(address.ID.Equals(50))
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

func TestJSONUnmarshallingRegisteredAddressFail(t *testing.T) {
	require := require.New(t)
	var address userent.RegisteredAddress

	// forçando teste do unmarshal
	err := address.UnmarshalJSON([]byte(`
		{
			"id":         50,	
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
	require.Equal(http.StatusInternalServerError, resperr.StatusCode(err))

	err = json.Unmarshal([]byte(`
	{
		"id":         -1,
		"Street":     "Street ABC",
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
	require.ErrorIs(err, uservo.ErrInvalidAddressID)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
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
	require.ErrorIs(err, uservo.ErrStreetMinLength)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
		"Street":     "Street ABC",
		"District":   "",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrDistrictMinLength)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
		"Street":     "Street ABC",
		"District":   "District",
		"City":       "",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrCityMinLength)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
		"Street":     "Street ABC",
		"District":   "District",
		"City":       "City",
		"State":      "",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrStateMinLength)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
		"Street":     "Street ABC",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrAddressNumberMinLength)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
		"Street":     "Street ABC",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "zipcodee",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrZipcodeNotNumeric)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
		"Street":     "Street ABC",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrLatitudeInvalidFormat)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
		"Street":     "Street ABC",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  ""
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrLongitudeInvalidFormat)

	err = json.Unmarshal([]byte(`
	{
		"id":         "",
		"Street":     "Street ABC",
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
	require.Equal(http.StatusBadRequest, resperr.StatusCode(err))

}
