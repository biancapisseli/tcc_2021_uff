package uservo

import (
	"fmt"

	"github.com/carlmjohnson/resperr"
	"github.com/google/uuid"

	"net/http"
)

type AddressID uuid.UUID

func (aid AddressID) String() string {
	return uuid.UUID(aid).String()
}

func (aid AddressID) Equals(other AddressID) bool {
	return uuid.UUID(aid).String() == uuid.UUID(other).String()
}

func NewAddressID(value string) (AddressID, error) {
	addressUUID, err := uuid.Parse(value)
	if err != nil || addressUUID == uuid.Nil {
		return AddressID(uuid.Nil), resperr.WithCodeAndMessage(
			fmt.Errorf("address id should be in valid UUID format: %w", err),
			http.StatusBadRequest,
			"o ID do endereço deve estar no formato de UUID",
		)
	}
	return AddressID(addressUUID), nil
}

func GenerateNewAddressID() (addressID AddressID) {
	return AddressID(uuid.New())
}
