package uservo

import (
	"fmt"

	"github.com/carlmjohnson/resperr"

	"net/http"
)

const (
	MaxRawPasswordLength = 30
	MinRawPasswordLength = 6
)

var (
	ErrRawPasswordMaxLength = fmt.Errorf("raw password should have < %d characters", MaxRawPasswordLength)
	ErrRawPasswordMinLength = fmt.Errorf("raw password should have > %d characters", MinRawPasswordLength)
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
			fmt.Sprintf("a senha deve ter no máximo %d caracteres", MaxRawPasswordLength),
		)
	}
	if len(value) < MinRawPasswordLength {
		return "", resperr.WithCodeAndMessage(
			ErrRawPasswordMinLength,
			http.StatusBadRequest,
			fmt.Sprintf("a senha deve ter no mínimo %d caracteres", MinRawPasswordLength),
		)
	}
	return PasswordRaw(value), nil
}
