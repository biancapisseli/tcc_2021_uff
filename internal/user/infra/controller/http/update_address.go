package userhttpctl

import (
	"context"
)

func (c UserHTTPController) UpdateAddress(
	ctx context.Context,
	userID userdom.UserID,
	address *userdom.RegisteredAddress,
) (err error) {
	return s.repo.SaveUserAddress(ctx, userID, address)
}
