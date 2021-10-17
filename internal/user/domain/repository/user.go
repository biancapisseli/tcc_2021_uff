package userrepo

import (
	"context"
	userent "ifoodish-store/internal/domain/valueobject"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

type UserRepository interface {
	GetUserByEmailAndPassword(ctx context.Context, email uservo.Email, password uservo.PasswordEncoded) (user *userent.RegisteredUser, err error)

	GetUserInfo(ctx context.Context, ID uservo.UserID) (user *userent.RegisteredUser, err error)
	AddUser(ctx context.Context, user *userent.User) (ID uservo.UserID, err error)
	SaveUser(ctx context.Context, user *userent.RegisteredUser) (err error)
	RemoveUser(ctx context.Context, ID uservo.UserID) (err error)
	UpdatePassword(ctx context.Context, ID uservo.UserID, newPassword uservo.PasswordEncoded) (err error)

	GetUserAddress(ctx context.Context, userID uservo.UserID, addresID uservo.AddressID) (address *userent.RegisteredAddress, err error)
	GetUserAddresses(ctx context.Context, userID uservo.UserID) (adresses []*userent.RegisteredAddress, err error)
	AddUserAddress(ctx context.Context, userID uservo.UserID, address *uservo.Address) (addresID uservo.AddressID, err error)
	SaveUserAddress(ctx context.Context, userID uservo.UserID, address *userent.RegisteredAddress) (err error)
	RemoveUserAddress(ctx context.Context, userID uservo.UserID, addresID uservo.AddressID) (err error)
}
