package uservo

import (
	"errors"
	"ifoodish-store/pkg/resperr"
	"net/http"
	"strconv"
)

const (
	MaxRawPasswordLength = 30
	MinRawPasswordLength = 6
)

var (
	ErrRawPasswordMaxLength = errors.New("passoword raw should have < " + strconv.Itoa(MaxRawPasswordLength) + " characters")
	ErrRawPasswordMinLength = errors.New("passoword raw should have > " + strconv.Itoa(MinRawPasswordLength) + " characters")
)

type PasswordRaw string

func (rp PasswordRaw) Equals(other PasswordRaw) bool {
	return rp.String() == other.String()
}

func (rp PasswordRaw) String() string {
	return string(rp)
}

func NewPasswordRaw(value string) (PasswordRaw, error) {
	if len(value) > MaxRawPasswordLength {
		return "", resperr.WithCodeAndMessage(
			ErrRawPasswordMaxLength,
			http.StatusBadRequest,
			"A senha está muito grande, deve ter menos que "+strconv.Itoa(MaxRawPasswordLength)+" digitos",
		)
	}
	if len(value) < MinRawPasswordLength {
		return "", resperr.WithCodeAndMessage(
			ErrRawPasswordMinLength,
			http.StatusBadRequest,
			"A senha está muito pequena, deve ter menos que "+strconv.Itoa(MinRawPasswordLength)+" digitos",
		)
	}
	return PasswordRaw(value), nil
}
