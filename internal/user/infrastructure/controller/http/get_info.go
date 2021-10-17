package userhttpctl

import (
	"context"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (c UserHTTPController) GetUserInfo(
	ctx context.Context,
	userID uservo.UserID,
	addressID uservo.AddressID,
) (userInfo *usernet.RegisteredUser, err error) {
	return c.repo.GetUserInfo(ctx, userID)
}
