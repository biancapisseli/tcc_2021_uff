package uservo

import "fmt"

var (
	ErrInvalidAddressID = fmt.Errorf("id do endereço é inválido")
	Invalid             = 0
)

type AddressID int64

func (uid AddressID) Equals(other AddressID) bool {
	return uid == other
}

func NewAddressID(value int64) (AddressID, error) {
	if value <= int64(Invalid) {
		return AddressID(Invalid), ErrInvalidAddressID
	}
	return AddressID(value), nil
}
