package usersvc

import (
	"context"
	userdom "ifoodish-store/user/domain"
)

func (s UserService) UpdateAddress(
	ctx context.Context,
	userID userdom.UserID,
	address *userdom.RegisteredAddress,
) (err error) {
	return s.repo.SaveUserAddress(ctx, userID, address)
}
