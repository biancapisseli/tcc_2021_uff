package uservo

import "fmt"

const (
	MaxRawPasswordLength = 30
	MinRawPasswordLength = 6
)

var (
	ErrRawPasswordMaxLength = fmt.Errorf("a senha deve possuir no máximo %d caracteres", MaxRawPasswordLength)
	ErrRawPasswordMinLength = fmt.Errorf("a senha deve possuir mais que %d caracteres", MinRawPasswordLength)
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
		return "", ErrRawPasswordMaxLength
	}
	if len(value) < MinRawPasswordLength {
		return "", ErrRawPasswordMinLength
	}
	return PasswordRaw(value), nil
}
