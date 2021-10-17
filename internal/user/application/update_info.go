package usersvc

import (
	"context"
	userent "ifoodish-store/internal/user/domain/entity"
)

func (s *UserService) UpdateInfo(
	ctx context.Context,
	user *userent.RegisteredUser,
) (err error) {
	return s.repo.SaveUser(ctx, user)
}