package userhttpctl

import (
	"context"
)

func (c UserHTTPController) AddAddress(
	ctx context.Context,
	userID userdom.UserID,
	address *userdom.Address,
) (addressID userdom.AddressID, err error) {
	return s.repo.AddUserAddress(ctx, userID, address)
}
