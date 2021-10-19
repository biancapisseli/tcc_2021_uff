package uservo

import (
	"errors"
	"ifoodish-store/pkg/resperr"
	"net/http"
	"strconv"
)

const (
	MaxPhoneLength = 11
	MinPhoneLength = 10
)

var (
	ErrPhoneMaxLength = errors.New("phone should have < " + strconv.Itoa(MaxPhoneLength) + " characteres")
	ErrPhoneMinLength = errors.New("phone should have > " + strconv.Itoa(MinPhoneLength) + " characteres")
)

type Phone string

func (p Phone) Equals(other Phone) bool {
	return p.String() == other.String()
}

func (p Phone) String() string {
	return string(p)
}

func NewPhone(value string) (Phone, error) {
	if len(value) > MaxPhoneLength {
		return "", resperr.WithCodeAndMessage(
			ErrPhoneMaxLength,
			http.StatusBadRequest,
			"O telefone está muito grande, deve ter menos que "+strconv.Itoa(MaxPhoneLength)+" digitos",
		)
	}
	if len(value) < MinPhoneLength {
		return "", resperr.WithCodeAndMessage(
			ErrPhoneMinLength,
			http.StatusBadRequest,
			"O telefone está muito grande, deve ter menos que "+strconv.Itoa(MinPhoneLength)+" digitos",
		)
	}
	return Phone(value), nil
}
