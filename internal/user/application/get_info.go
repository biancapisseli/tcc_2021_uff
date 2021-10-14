package usersvc

import (
	"context"
	userent "ifoodish-store/internal/domain/entity"
)

func (s UserService) GetUserInfo(
	ctx context.Context,
	userID userent.UserID,
	addressID userent.AddressID,
) (userInfo *userent.RegisteredUser, err error) {
	return s.repo.GetUserInfo(ctx, userID)
}
