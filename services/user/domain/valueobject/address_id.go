package uservo

import (
	"errors"
	"fmt"

	"github.com/carlmjohnson/resperr"

	"net/http"
)

var (
	ErrInvalidAddressID = errors.New("address ID should be numeric and > 0")
)

type AddressID int64

func (aid AddressID) String() string {
	return fmt.Sprintf("%d", int64(aid))
}

func (aid AddressID) Equals(other AddressID) bool {
	return aid.String() == other.String()
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
