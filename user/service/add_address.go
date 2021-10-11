package usersvc

import (
	"context"
	userdom "ifoodish-store/user/domain"
)

func (s UserService) AddAddress(
	ctx context.Context,
	userID userdom.UserID,
	address *userdom.Address,
) (addressID userdom.AddressID, err error) {
	return s.repo.AddUserAddress(ctx, userID, address)
}
