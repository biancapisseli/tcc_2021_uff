package useruc

import (
	"context"
	"fmt"
	userent "ifoodish-store/services/user/domain/entity"
)

func (s UserUseCases) UpdateUserInfo(
	ctx context.Context,
	user userent.RegisteredUser,
) (err error) {
	err = s.repo.SaveUser(ctx, user)
	if err != nil {
		return fmt.Errorf("error updating user info: %w", err)
	}
	return nil
}
