package useruc

import (
	"context"
	userent "ifoodish-store/internal/user/domain/entity"
)

func (s UserUseCases) UpdateUserInfo(
	ctx context.Context,
	user userent.RegisteredUser,
) (err error) {
	return s.repo.SaveUser(ctx, user)
}
