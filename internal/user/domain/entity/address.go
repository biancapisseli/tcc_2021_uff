package userent

import (
	"encoding/json"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

type RegisteredAddress struct {
	Address
	ID uservo.AddressID `json:"id"`
}

type Address struct {
	Street     uservo.Street     `json:"street"`
	District   uservo.District   `json:"district"`
	City       uservo.City       `json:"city"`
	State      uservo.State      `json:"state"`
	Complement uservo.Complement `json:"complement"`
	Number     uservo.Number     `json:"number"`
	Zipcode    uservo.Zipcode    `json:"zipcode"`
	Latitude   uservo.Latitude   `json:"latitude"`
	Longitude  uservo.Longitude  `json:"longitude"`
}

func NewRegisteredAddress(params RegisteredAddress) (newAddress *RegisteredAddress, err error) {
	newAddress = new(RegisteredAddress)
	childAddress, err := NewAddress(params.Address)
	if err != nil {
		return nil, err
	}
	newAddress.Address = *childAddress

	newAddress.ID, err = uservo.NewAddressID(int64(params.ID))
	if err != nil {
		return nil, err
	}
	return newAddress, nil
}

func NewAddress(address Address) (newAddress *Address, err error) {
	newAddress = new(Address)
	newAddress.Street, err = uservo.NewStreet(string(address.Street))
	if err != nil {
		return nil, err
	}
	newAddress.District, err = uservo.NewDistrict(string(address.District))
	if err != nil {
		return nil, err
	}
	newAddress.City, err = uservo.NewCity(string(address.City))
	if err != nil {
		return nil, err
	}
	newAddress.State, err = uservo.NewState(string(address.State))
	if err != nil {
		return nil, err
	}
	newAddress.Complement, err = uservo.NewComplement(string(address.Complement))
	if err != nil {
		return nil, err
	}
	newAddress.Number, err = uservo.NewAddressNumber(string(address.Number))
	if err != nil {
		return nil, err
	}
	newAddress.Zipcode, err = uservo.NewZipcode(string(address.Zipcode))
	if err != nil {
		return nil, err
	}
	newAddress.Latitude, err = uservo.NewLatitude(string(address.Latitude))
	if err != nil {
		return nil, err
	}
	newAddress.Longitude, err = uservo.NewLongitude(string(address.Longitude))
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
