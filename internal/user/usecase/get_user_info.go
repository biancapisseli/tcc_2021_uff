package useruc

import (
	"context"
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (s UserUseCases) GetUserInfo(
	ctx context.Context,
	userID uservo.UserID,
) (userInfo userent.RegisteredUser, err error) {
	userInfo, err = s.repo.GetUserInfo(ctx, userID)
	if err != nil {
		return userInfo, fmt.Errorf("error gettin user info: %w", err)
	}
	return userInfo, nil
}
