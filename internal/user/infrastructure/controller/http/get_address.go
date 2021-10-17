package userhttpctl

import (
	"context"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (c UserHTTPController) GetAddress(
	ctx context.Context,
	userID uservo.UserID,
	addressID uservo.AddressID,
) (address *userent.RegisteredAddress, err error) {
	return c.repo.GetUserAddress(ctx, userID, addressID)
}
