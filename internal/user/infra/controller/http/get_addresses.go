package userhttpctl

import (
	"context"
)

func (c UserHTTPController) GetAddresses(
	ctx context.Context,
	userID userdom.UserID,
	addressID userdom.AddressID,
) (addresses []*userdom.RegisteredAddress, err error) {
	return s.repo.GetUserAddresses(ctx, userID)
}
