package usersvc

import (
	"context"
	userdom "ifoodish-store/user/domain"
)

func (s *UserService) UpdateInfo(
	ctx context.Context,
	user *userdom.RegisteredUser,
) (err error) {
	return s.repo.SaveUser(ctx, user)
}
