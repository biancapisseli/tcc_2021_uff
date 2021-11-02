package userent

import (
	"encoding/json"
	"fmt"
	uservo "ifoodish-store/services/user/domain/valueobject"
	"net/http"

	"github.com/carlmjohnson/resperr"
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

func NewRegisteredUser(id string, user User) (newUser RegisteredUser, err error) {
	newUser.ID, err = uservo.NewUserID(id)
	if err != nil {
		return newUser, fmt.Errorf("error creating new registered user id: %w", err)
	}
	newUser.User = user
	return newUser, nil
}

func NewUser(name, email, phone string) (newUser User, err error) {
	newUser.Name, err = uservo.NewUserName(name)
	if err != nil {
		return newUser, fmt.Errorf("error creating new user name: %w", err)
	}
	newUser.Email, err = uservo.NewEmail(email)
	if err != nil {
		return newUser, fmt.Errorf("error creating new user email: %w", err)
	}
	newUser.Phone, err = uservo.NewPhone(phone)
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

	newUser, err := NewUser(
		userClone.Name.String(),
		userClone.Email.String(),
		userClone.Phone.String(),
	)
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
		UserID string `json:"id"`
	}
	if err := json.Unmarshal(data, &registered); err != nil {
		return resperr.WithStatusCode(
			fmt.Errorf("error unmarshalling registered user: %w", err),
			http.StatusBadRequest,
		)
	}

	newRegisteredUser, err := NewRegisteredUser(registered.UserID, user)
	if err != nil {
		return fmt.Errorf("error unmarshalling registered user: %w", err)
	}

	*u = newRegisteredUser
	return nil
}
