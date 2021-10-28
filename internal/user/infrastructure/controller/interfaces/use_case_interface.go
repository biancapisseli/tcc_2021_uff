package userctlint

import (
	"context"

	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

type UserUseCases interface {
	AddUserAddress(
		ctx context.Context,
		userID uservo.UserID,
		address userent.Address,
	) (addressID uservo.AddressID, err error)

	ChangePassword(
		ctx context.Context,
		userID uservo.UserID,
		currentPassword uservo.PasswordRaw,
		newPassword uservo.PasswordRaw,
		newPasswordConfirm uservo.PasswordRaw,
	) (err error)

	GetUserAddress(
		ctx context.Context,
		userID uservo.UserID,
		addressID uservo.AddressID,
	) (address userent.RegisteredAddress, err error)

	GetUserAddresses(
		ctx context.Context,
		userID uservo.UserID,
	) (addresses []userent.RegisteredAddress, err error)

	GetUserInfo(
		ctx context.Context,
		userID uservo.UserID,
	) (userInfo userent.RegisteredUser, err error)

	RegisterUser(
		ctx context.Context,
		user userent.User,
		password uservo.PasswordRaw,
		passwordConfirm uservo.PasswordRaw,
	) (userID uservo.UserID, err error)

	RemoveUserAddress(
		ctx context.Context,
		userID uservo.UserID,
		addressID uservo.AddressID,
	) (err error)

	UpdateUserAddress(
		ctx context.Context,
		userID uservo.UserID,
		address userent.RegisteredAddress,
	) (err error)

	UpdateUserInfo(
		ctx context.Context,
		user userent.RegisteredUser,
	) (err error)
}
