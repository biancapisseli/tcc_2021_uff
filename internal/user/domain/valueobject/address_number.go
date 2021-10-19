package uservo

import (
	"errors"
	"ifoodish-store/pkg/resperr"
	"net/http"
	"strconv"
)

const (
	MaxAddressNumberLength = 10
	MinAddressNumberLength = 1
)

var (
	ErrAddressNumberMaxLength = errors.New(" address_number should have < " + strconv.Itoa(MaxAddressNumberLength) + " characteres")
	ErrAddressNumberMinLength = errors.New(" address_number should have > " + strconv.Itoa(MinAddressNumberLength) + " characteres")
)

type AddressNumber string

func (s AddressNumber) Equals(other AddressNumber) bool {
	return s.String() == other.String()
}

func (s AddressNumber) String() string {
	return string(s)
}

func NewAddressNumber(value string) (AddressNumber, error) {
	if len(value) > MaxAddressNumberLength {
		return "", resperr.WithCodeAndMessage(
			ErrAddressNumberMaxLength,
			http.StatusBadRequest,
			"O numero está muito grande, deve ter menos que "+strconv.Itoa(MaxAddressNumberLength)+"digitos",
		)
	}
	if len(value) <= MinAddressNumberLength {
		return "", resperr.WithCodeAndMessage(
			ErrAddressNumberMinLength,
			http.StatusBadRequest,
			"O numero está muito pequeno, deve ter mais que "+strconv.Itoa(MinAddressNumberLength)+"digito",
		)
	}

	return AddressNumber(value), nil
}
