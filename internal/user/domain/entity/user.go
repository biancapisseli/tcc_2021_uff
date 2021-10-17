package userent

import (
	"encoding/json"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

type RegisteredUser struct {
	User
	ID uservo.UserID `json:"id"`
}

type User struct {
	Name  uservo.UserName `json:"name"`
	Email uservo.Email    `json:"email"`
	Phone uservo.Phone    `json:"phone"`
}

func NewRegisteredUser(params RegisteredUser) (newUser *RegisteredUser, err error) {
	newUser = new(RegisteredUser)
	childUser, err := NewUser(params.User)
	if err != nil {
		return nil, err
	}
	newUser.User = *childUser

	newUser.ID, err = uservo.NewUserID(int64(params.ID))
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func NewUser(params User) (newUser *User, err error) {
	newUser = new(User)
	newUser.Name, err = uservo.NewUserName(string(params.Name))
	if err != nil {
		return nil, err
	}
	newUser.Email, err = uservo.NewEmail(string(params.Email))
	if err != nil {
		return nil, err
	}
	newUser.Phone, err = uservo.NewPhone(string(params.Phone))
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *User) UnmarshalJSON(data []byte) error {

	type clone User
	var userClone clone

	if err := json.Unmarshal(data, &userClone); err != nil {
		return err
	}

	newUser, err := NewUser(User(userClone))
	if err != nil {
		return err
	}

	*u = *newUser
	return nil
}
