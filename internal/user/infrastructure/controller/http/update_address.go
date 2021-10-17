package userhttpctl

import (
	"context"
	userent "ifoodish-store/internal/user/domain/entity"

	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (c UserHTTPController) UpdateAddress(
	ctx context.Context,
	userID uservo.UserID,
	address *userent.RegisteredAddress,
) (err error) {
	return c.repo.SaveUserAddress(ctx, userID, address)
}
