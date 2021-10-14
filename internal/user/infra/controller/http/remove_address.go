package userhttpctl

import (
	"context"
)

func (c UserHTTPController) RemoveAddress(
	ctx context.Context,
	userID userdom.UserID,
	addressID userdom.AddressID,
) (err error) {
	return s.repo.RemoveUserAddress(ctx, userID, addressID)
}
