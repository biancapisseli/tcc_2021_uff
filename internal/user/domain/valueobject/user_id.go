package uservo

import (
	"errors"
	"ifoodish-store/pkg/resperr"
	"net/http"
)

var (
	ErrInvalidUserID = errors.New("user ID shoul be > 0")
)

type UserID int64

func (uid UserID) Equals(other UserID) bool {
	return uid == other
}

func NewUserID(value int64) (UserID, error) {
	if value <= 0 {
		return 0, resperr.WithCodeAndMessage(
			ErrInvalidUserID,
			http.StatusBadRequest,
			"O ID do usuário deve ser maior que 0",
		)
	}
	return UserID(value), nil
}
