package uservo

import "fmt"

var (
	ErrInvalidUserID = fmt.Errorf("id do usuário é inválido")
)

type UserID int64

func (uid UserID) Equals(other UserID) bool {
	return uid == other
}

func NewUserID(value int64) (UserID, error) {
	if value <= 0 {
		return 0, ErrInvalidUserID
	}
	return UserID(value), nil
}
