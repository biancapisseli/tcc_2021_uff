package userhttpctl

import (
	"context"
)

func (s *UserService) UpdateInfo(
	ctx context.Context,
	user *userdom.RegisteredUser,
) (err error) {
	return s.repo.SaveUser(ctx, user)
}
