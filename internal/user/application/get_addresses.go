package usersvc

import (
	"context"
	userent "ifoodish-store/internal/domain/entity"
)

func (s UserService) GetAddresses(
	ctx context.Context,
	userID userent.UserID,
	addressID userent.AddressID,
) (addresses []*userent.RegisteredAddress, err error) {
	return s.repo.GetUserAddresses(ctx, userID)
}
