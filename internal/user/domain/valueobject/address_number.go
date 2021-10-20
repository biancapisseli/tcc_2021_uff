package uservo

import (
	"fmt"
	"ifoodish-store/pkg/resperr"

	"net/http"
)

const (
	MaxAddressNumberLength = 10
	MinAddressNumberLength = 1
)

var (
	ErrAddressNumberMaxLength = fmt.Errorf("address_number should have < %d characters", MaxAddressNumberLength)
	ErrAddressNumberMinLength = fmt.Errorf("address_number should have > %d characters", MinAddressNumberLength)
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
			fmt.Sprintf("o número do endereço deve ter no máximo %d caracteres", MaxAddressNumberLength),
		)
	}
	if len(value) <= MinAddressNumberLength {
		return "", resperr.WithCodeAndMessage(
			ErrAddressNumberMinLength,
			http.StatusBadRequest,
			fmt.Sprintf("o número do endereço deve ter no máximo %d caracteres", MaxAddressNumberLength),
		)
	}

	return AddressNumber(value), nil
}
