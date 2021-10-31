package uservo

import (
	"fmt"

	"github.com/carlmjohnson/resperr"

	"net/http"

	"github.com/google/uuid"
)

type UserID uuid.UUID

func (uid UserID) Equals(other UserID) bool {
	return uuid.UUID(uid).String() == uuid.UUID(other).String()
}

func (uid UserID) String() string {
	return uuid.UUID(uid).String()
}

func NewUserID(value string) (UserID, error) {
	userUUID, err := uuid.Parse(value)
	if err != nil || userUUID == uuid.Nil {
		return UserID(uuid.Nil), resperr.WithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"o ID do usuário deve estar no formato de UUID",
		)
	}

	return UserID(userUUID), nil
}

func GenerateNewUserID() (userID UserID, err error) {
	userUUID, err := uuid.NewRandom()
	if err != nil {
		return userID, resperr.WithStatusCode(
			fmt.Errorf("error generating new user id: %w", err),
			http.StatusInternalServerError,
		)
	}

	return UserID(userUUID), nil
}
