package userrepo

import (
	"context"
	uservo "ifoodish-store/internal/domain/valueobject"
)

type UserRepository interface {
	GetUserByEmailAndPassword(ctx context.Context, email uservo.Email, password uservo.PasswordEncoded) (user *uservo.RegisteredUser, err error)

	GetUserInfo(ctx context.Context, ID uservo.UserID) (user *uservo.RegisteredUser, err error)
	AddUser(ctx context.Context, user *uservo.User) (ID uservo.UserID, err error)
	SaveUser(ctx context.Context, user *uservo.RegisteredUser) (err error)
	RemoveUser(ctx context.Context, ID uservo.UserID) (err error)
	UpdatePassword(ctx context.Context, ID uservo.UserID, newPassword uservo.PasswordEncoded) (err error)

	GetUserAddress(ctx context.Context, userID uservo.UserID, addresID uservo.AddressID) (address *uservo.RegisteredAddress, err error)
	GetUserAddresses(ctx context.Context, userID uservo.UserID) (adresses []*uservo.RegisteredAddress, err error)
	AddUserAddress(ctx context.Context, userID uservo.UserID, address *uservo.Address) (addresID uservo.AddressID, err error)
	SaveUserAddress(ctx context.Context, userID uservo.UserID, address *uservo.RegisteredAddress) (err error)
	RemoveUserAddress(ctx context.Context, userID uservo.UserID, addresID uservo.AddressID) (err error)
}
