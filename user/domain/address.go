package userdom

import "encoding/json"

type RegisteredAddress struct {
	Address
	ID AddressID `json:"id"`
}

type Address struct {
	Street     Street     `json:"street"`
	District   District   `json:"district"`
	City       City       `json:"city"`
	State      State      `json:"state"`
	Complement Complement `json:"complement"`
	Number     Number     `json:"number"`
	Zipcode    Zipcode    `json:"zipcode"`
	Latitude   Latitude   `json:"latitude"`
	Longitude  Longitude  `json:"longitude"`
}

func NewRegisteredAddress(params RegisteredAddress) (newAddress *RegisteredAddress, err error) {
	newAddress = new(RegisteredAddress)
	childAddress, err := NewAddress(params.Address)
	if err != nil {
		return nil, err
	}
	newAddress.Address = *childAddress

	newAddress.ID, err = NewAddressID(int64(params.ID))
	if err != nil {
		return nil, err
	}
	return newAddress, nil
}

func NewAddress(address Address) (newAddress *Address, err error) {
	newAddress = new(Address)
	newAddress.Street, err = NewStreet(string(address.Street))
	if err != nil {
		return nil, err
	}
	newAddress.District, err = NewDistrict(string(address.District))
	if err != nil {
		return nil, err
	}
	newAddress.City, err = NewCity(string(address.City))
	if err != nil {
		return nil, err
	}
	newAddress.State, err = NewState(string(address.State))
	if err != nil {
		return nil, err
	}
	newAddress.Complement, err = NewComplement(string(address.Complement))
	if err != nil {
		return nil, err
	}
	newAddress.Number, err = NewNumber(string(address.Number))
	if err != nil {
		return nil, err
	}
	newAddress.Zipcode, err = NewZipcode(string(address.Zipcode))
	if err != nil {
		return nil, err
	}
	newAddress.Latitude, err = NewLatitude(string(address.Latitude))
	if err != nil {
		return nil, err
	}
	newAddress.Longitude, err = NewLongitude(string(address.Longitude))
	if err != nil {
		return nil, err
	}
	return newAddress, nil
}

func (a *Address) UnmarshalJSON(data []byte) error {

	type clone Address
	var addressClone clone

	if err := json.Unmarshal(data, &addressClone); err != nil {
		return err
	}

	newAddress, err := NewAddress(Address(addressClone))
	if err != nil {
		return err
	}

	*a = *newAddress

	return nil
}
