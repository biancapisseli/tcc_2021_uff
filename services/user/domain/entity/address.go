package userent

import (
	"encoding/json"
	"fmt"
	uservo "ifoodish-store/services/user/domain/valueobject"
	"net/http"

	"github.com/carlmjohnson/resperr"
)

type RegisteredAddress struct {
	Address
	ID uservo.AddressID `json:"id"`
}

type Address struct {
	Street     uservo.Street        `json:"street"`
	District   uservo.District      `json:"district"`
	City       uservo.City          `json:"city"`
	State      uservo.State         `json:"state"`
	Complement uservo.Complement    `json:"complement"`
	Number     uservo.AddressNumber `json:"number"`
	Zipcode    uservo.Zipcode       `json:"zipcode"`
	Latitude   uservo.Latitude      `json:"latitude"`
	Longitude  uservo.Longitude     `json:"longitude"`
}

func NewRegisteredAddress(id int64, address Address) (newAddress RegisteredAddress, err error) {
	newAddress.ID, err = uservo.NewAddressID(id)
	if err != nil {
		return newAddress, fmt.Errorf("error creating new registered address id: %w", err)
	}
	newAddress.Address = address
	return newAddress, nil
}

func NewAddress(
	street, district, city, state, complement,
	number, zipcode, latitude, longitude string,
) (newAddress Address, err error) {
	newAddress.Street, err = uservo.NewStreet(string(street))
	if err != nil {
		return newAddress, fmt.Errorf(
			"error creating new address street: %w",
			err,
		)
	}
	newAddress.District, err = uservo.NewDistrict(district)
	if err != nil {
		return newAddress, fmt.Errorf(
			"error creating new address district: %w",
			err,
		)
	}
	newAddress.City, err = uservo.NewCity(city)
	if err != nil {
		return newAddress, fmt.Errorf(
			"error creating new address city: %w",
			err,
		)
	}
	newAddress.State, err = uservo.NewState(state)
	if err != nil {
		return newAddress, fmt.Errorf(
			"error creating new address state: %w",
			err,
		)
	}
	newAddress.Complement, err = uservo.NewComplement(complement)
	if err != nil {
		return newAddress, fmt.Errorf(
			"error creating new address complement: %w",
			err,
		)
	}
	newAddress.Number, err = uservo.NewAddressNumber(number)
	if err != nil {
		return newAddress, fmt.Errorf(
			"error creating new address number: %w",
			err,
		)
	}
	newAddress.Zipcode, err = uservo.NewZipcode(zipcode)
	if err != nil {
		return newAddress, fmt.Errorf(
			"error creating new address zipcode: %w",
			err,
		)
	}
	newAddress.Latitude, err = uservo.NewLatitude(latitude)
	if err != nil {
		return newAddress, fmt.Errorf(
			"error creating new address latitude: %w",
			err,
		)
	}
	newAddress.Longitude, err = uservo.NewLongitude(longitude)
	if err != nil {
		return newAddress, fmt.Errorf(
			"error creating new address longitude: %w",
			err,
		)
	}
	return newAddress, nil
}

func (a *Address) UnmarshalJSON(data []byte) error {

	type clone Address
	var addressClone clone

	if err := json.Unmarshal(data, &addressClone); err != nil {
		return fmt.Errorf("error unmarshalling address: %w", err)
	}

	newAddress, err := NewAddress(
		addressClone.Street.String(),
		addressClone.District.String(),
		addressClone.City.String(),
		addressClone.State.String(),
		addressClone.Complement.String(),
		addressClone.Number.String(),
		addressClone.Zipcode.String(),
		addressClone.Latitude.String(),
		addressClone.Longitude.String(),
	)
	if err != nil {
		return fmt.Errorf("error unmarshalling address: %w", err)
	}

	*a = newAddress

	return nil
}

func (u *RegisteredAddress) UnmarshalJSON(data []byte) error {

	var address Address
	if err := json.Unmarshal(data, &address); err != nil {
		return fmt.Errorf("error unmarshalling registered address: %w", err)
	}

	var registered struct {
		AddressID int64 `json:"id"`
	}

	if err := json.Unmarshal(data, &registered); err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf("error unmarshalling registered address: %w", err),
			http.StatusBadRequest,
		)
	}

	newRegisteredAddress, err := NewRegisteredAddress(registered.AddressID, address)
	if err != nil {
		return fmt.Errorf("error unmarshalling registered address: %w", err)
	}

	*u = newRegisteredAddress
	return nil
}
