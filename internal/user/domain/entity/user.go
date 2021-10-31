package userent

import (
	"encoding/json"
	"fmt"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"net/http"

	"github.com/carlmjohnson/resperr"

	"github.com/google/uuid"
)

type RegisteredUser struct {
	ID uservo.UserID `json:"id"`
	User
}

type User struct {
	Name  uservo.UserName `json:"name"`
	Email uservo.Email    `json:"email"`
	Phone uservo.Phone    `json:"phone"`
}

func NewRegisteredUser(params RegisteredUser) (newUser RegisteredUser, err error) {
	newUser.User, err = NewUser(params.User)
	if err != nil {
		return newUser, fmt.Errorf("error creating new registered user: %w", err)
	}

	newUser.ID, err = uservo.NewUserID(params.ID.String())
	if err != nil {
		return newUser, fmt.Errorf("error creating new registered user id: %w", err)
	}
	return newUser, nil
}

func NewUser(params User) (newUser User, err error) {
	newUser.Name, err = uservo.NewUserName(string(params.Name))
	if err != nil {
		return newUser, fmt.Errorf("error creating new user name: %w", err)
	}
	newUser.Email, err = uservo.NewEmail(string(params.Email))
	if err != nil {
		return newUser, fmt.Errorf("error creating new user email: %w", err)
	}
	newUser.Phone, err = uservo.NewPhone(string(params.Phone))
	if err != nil {
		return newUser, fmt.Errorf("error creating new user phone: %w", err)
	}
	return newUser, nil
}

func (u *User) UnmarshalJSON(data []byte) error {

	type clone User
	var userClone clone

	if err := json.Unmarshal(data, &userClone); err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf("error unmarshalling user: %w", err),
			http.StatusBadRequest,
		)
	}

	newUser, err := NewUser(User(userClone))
	if err != nil {
		return fmt.Errorf("error unmarshalling user: %w", err)
	}

	*u = newUser
	return nil
}

func (u *RegisteredUser) UnmarshalJSON(data []byte) error {

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		return fmt.Errorf("error unmarshalling registered user: %w", err)
	}

	var registered struct {
		UserID uuid.UUID `json:"id"`
	}

	if err := json.Unmarshal(data, &registered); err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf("error unmarshalling registered user: %w", err),
			http.StatusBadRequest,
		)
	}

	newRegisteredUser, err := NewRegisteredUser(RegisteredUser{
		ID:   uservo.UserID(registered.UserID),
		User: user,
	})
	if err != nil {
		return fmt.Errorf("error unmarshalling registered user: %w", err)
	}

	*u = newRegisteredUser
	return nil
}
