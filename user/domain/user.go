package userdom

import (
	"context"
	"encoding/json"
)

type RegisteredUser struct {
	User
	ID UserID `json:"id"`
}

type User struct {
	Name  UserName `json:"name"`
	Email Email    `json:"email"`
}

type UserRepository interface {
	GetUserInfo(ctx context.Context, ID UserID) (user *RegisteredUser, err error)
	AddUser(ctx context.Context, user *User) (ID UserID, err error)
	SaveUser(ctx context.Context, user *RegisteredUser) (err error)
	RemoveUser(ctx context.Context, ID UserID) (err error)
	UpdatePassword(ctx context.Context, ID UserID, newPassword PasswordEncoded) (err error)

	GetUserAddress(ctx context.Context, userID UserID, addresID AddressID) (address *RegisteredAddress, err error)
	GetUserAddresses(ctx context.Context, userID UserID) (adresses []*RegisteredAddress, err error)
	AddUserAddress(ctx context.Context, userID UserID, address *Address) (addresID AddressID, err error)
	SaveUserAddress(ctx context.Context, userID UserID, address *RegisteredAddress) (err error)
	RemoveUserAddress(ctx context.Context, userID UserID, addresID AddressID) (err error)
}

type PasswordEncoder interface {
	EncodePassword(rawPassword PasswordRaw) (encodedPassword PasswordEncoded, err error)
}

func NewRegisteredUser(params RegisteredUser) (newUser *RegisteredUser, err error) {
	newUser = new(RegisteredUser)
	childUser, err := NewUser(params.User)
	if err != nil {
		return nil, err
	}
	newUser.User = *childUser

	newUser.ID, err = NewUserID(int64(params.ID))
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func NewUser(params User) (newUser *User, err error) {
	newUser = new(User)
	newUser.Name, err = NewUserName(string(params.Name))
	if err != nil {
		return nil, err
	}
	newUser.Email, err = NewEmail(string(params.Email))
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
