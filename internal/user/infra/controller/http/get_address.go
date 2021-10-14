package userhttpctl

import (
	"context"
)

func (c UserHTTPController) GetAddress(
	ctx context.Context,
	userID userdom.UserID,
	addressID userdom.AddressID,
) (address *userdom.RegisteredAddress, err error) {
	return s.repo.GetUserAddress(ctx, userID, addressID)
}
