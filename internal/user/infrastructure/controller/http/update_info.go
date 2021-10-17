package userhttpctl

import (
	"context"
	userent "ifoodish-store/internal/domain/valueobject"
)

func (r UserHTTPController) UpdateInfo(
	ctx context.Context,
	user *userent.RegisteredUser,
) (err error) {
	return c.repo.SaveUser(ctx, user)
}
