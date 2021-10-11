package usersvc

import (
	"context"
	userdom "ifoodish-store/user/domain"
)

func (s UserService) RemoveAddress(
	ctx context.Context,
	userID userdom.UserID,
	addressID userdom.AddressID,
) (err error) {
	return s.repo.RemoveUserAddress(ctx, userID, addressID)
}
