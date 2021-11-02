package useruc

import (
	"context"
	"fmt"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
)

func (s UserUseCases) UpdateUserInfo(
	ctx context.Context,
	userID uservo.UserID,
	user userent.User,
) (err error) {
	err = s.repo.SaveUser(ctx, userID, user)
	if err != nil {
		return fmt.Errorf("error updating user info: %w", err)
	}
	return nil
}
