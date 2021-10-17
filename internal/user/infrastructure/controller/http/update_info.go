package userhttpctl

import (
	"context"
	userent "ifoodish-store/internal/user/domain/entity"
)

func (c UserHTTPController) UpdateInfo(
	ctx context.Context,
	user *userent.RegisteredUser,
) (err error) {
	return c.repo.SaveUser(ctx, user)
}
