package usersvc

import (
	"context"
	userdom "ifoodish-store/user/domain"
)

func (s UserService) GetAddress(
	ctx context.Context,
	userID userdom.UserID,
	addressID userdom.AddressID,
) (address *userdom.RegisteredAddress, err error) {
	return s.repo.GetUserAddress(ctx, userID, addressID)
}
