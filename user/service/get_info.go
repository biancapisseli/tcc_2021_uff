package usersvc

import (
	"context"
	userdom "ifoodish-store/user/domain"
)

func (s UserService) GetUserInfo(
	ctx context.Context,
	userID userdom.UserID,
	addressID userdom.AddressID,
) (userInfo *userdom.RegisteredUser, err error) {
	return s.repo.GetUserInfo(ctx, userID)
}
