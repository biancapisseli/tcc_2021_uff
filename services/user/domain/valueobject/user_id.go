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
			fmt.Errorf("user id should be in valid UUID format: %w", err),
			http.StatusBadRequest,
			"o ID do usuário deve estar no formato de UUID",
		)
	}

	return UserID(userUUID), nil
}

func GenerateNewUserID() (userID UserID) {
	return UserID(uuid.New())
}
