package userhttpctl

import (
	"context"
	userent "ifoodish-store/internal/domain/valueobject"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (c UserHTTPController) AddAddress(
	ctx context.Context,
	userID uservo.UserID,
	address *userent.Address,
) (addressID uservo.AddressID, err error) {
	return c.repo.AddUserAddress(ctx, userID, address)
}
