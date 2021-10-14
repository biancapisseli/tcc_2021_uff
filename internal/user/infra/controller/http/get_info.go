package userhttpctl

import (
	"context"
)

func (c UserHTTPController) GetUserInfo(
	ctx context.Context,
	userID userdom.UserID,
	addressID userdom.AddressID,
) (userInfo *userdom.RegisteredUser, err error) {
	return s.repo.GetUserInfo(ctx, userID)
}
