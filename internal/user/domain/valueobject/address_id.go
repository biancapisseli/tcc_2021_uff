package uservo

import (
	"errors"
	"ifoodish-store/pkg/resperr"

	"net/http"
)

var (
	ErrInvalidAddressID = errors.New("address ID should be numeric and > 0")
)

type AddressID int64

func (uid AddressID) Equals(other AddressID) bool {
	return uid == other
}

func NewAddressID(value int64) (AddressID, error) {
	if value <= int64(0) {
		return AddressID(0), resperr.WithCodeAndMessage(
			ErrInvalidAddressID,
			http.StatusBadRequest,
			"o ID do endereço deve ser maior que zero",
		)
	}
	return AddressID(value), nil
}
