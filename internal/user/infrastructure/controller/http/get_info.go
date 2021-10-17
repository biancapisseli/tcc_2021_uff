package userhttpctl

import (
	"context"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (c UserHTTPController) GetUserInfo(
	ctx context.Context,
	userID uservo.UserID,
	addressID uservo.AddressID,
) (userInfo *userent.RegisteredUser, err error) {
	return c.repo.GetUserInfo(ctx, userID)
}
