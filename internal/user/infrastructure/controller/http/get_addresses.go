package userhttpctl

import (
	"context"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (c UserHTTPController) GetAddresses(
	ctx context.Context,
	userID uservo.UserID,
	addressID uservo.AddressID,
) (addresses []*userent.RegisteredAddress, err error) {
	return c.repo.GetUserAddresses(ctx, userID)
}
