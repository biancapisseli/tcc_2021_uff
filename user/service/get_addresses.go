package usersvc

import (
	"context"
	userdom "ifoodish-store/user/domain"
)

func (s UserService) GetAddresses(
	ctx context.Context,
	userID userdom.UserID,
	addressID userdom.AddressID,
) (addresses []*userdom.RegisteredAddress, err error) {
	return s.repo.GetUserAddresses(ctx, userID)
}
