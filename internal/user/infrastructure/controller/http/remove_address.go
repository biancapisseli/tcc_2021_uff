package userhttpctl

import (
	"context"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (c UserHTTPController) RemoveAddress(
	ctx context.Context,
	userID uservo.UserID,
	addressID uservo.AddressID,
) (err error) {
	return c.repo.RemoveUserAddress(ctx, userID, addressID)
}
