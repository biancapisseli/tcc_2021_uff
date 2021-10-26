package userrepo

import (
	"context"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

type UserRepository interface {
	GetUserByEmailAndPassword(ctx context.Context, email uservo.Email, password uservo.PasswordEncoded) (user userent.RegisteredUser, err error)

	GetUserInfo(ctx context.Context, userID uservo.UserID) (user userent.RegisteredUser, err error)
	AddUser(ctx context.Context, user userent.User, password uservo.PasswordEncoded) (userID uservo.UserID, err error)
	SaveUser(ctx context.Context, user userent.RegisteredUser) (err error)
	RemoveUser(ctx context.Context, userID uservo.UserID) (err error)
	UpdatePassword(ctx context.Context, userID uservo.UserID, newPassword uservo.PasswordEncoded) (err error)

	GetUserAddress(ctx context.Context, userID uservo.UserID, addressID uservo.AddressID) (address userent.RegisteredAddress, err error)
	GetUserAddresses(ctx context.Context, userID uservo.UserID) (adresses []userent.RegisteredAddress, err error)
	AddUserAddress(ctx context.Context, userID uservo.UserID, address userent.Address) (addressID uservo.AddressID, err error)
	SaveUserAddress(ctx context.Context, userID uservo.UserID, address userent.RegisteredAddress) (err error)
	RemoveUserAddress(ctx context.Context, userID uservo.UserID, addressID uservo.AddressID) (err error)
}
