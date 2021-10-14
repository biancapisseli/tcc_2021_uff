package uservo

import "fmt"

const (
	maxRawPasswordLength = 30
	minRawPasswordLength = 6
)

var (
	ErrRawPasswordMaxLength = fmt.Errorf("a senha deve possuir menos que %d caracteres", maxRawPasswordLength)
	ErrRawPasswordMinLength = fmt.Errorf("a senha deve possuir mais que %d caracteres", minRawPasswordLength)
)

type PasswordRaw string

func (rp PasswordRaw) Equals(other PasswordRaw) bool {
	return rp.String() == other.String()
}

func (rp PasswordRaw) String() string {
	return string(rp)
}

func NewPasswordRaw(value string) (PasswordRaw, error) {
	if len(value) > maxRawPasswordLength {
		return "", ErrRawPasswordMaxLength
	}
	if len(value) < minRawPasswordLength {
		return "", ErrRawPasswordMinLength
	}
	return PasswordRaw(value), nil
}
