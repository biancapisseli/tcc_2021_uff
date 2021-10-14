package usersvc

import (
	"context"
	userent "ifoodish-store/internal/domain/entity"
)

func (s UserService) GetAddress(
	ctx context.Context,
	userID userent.UserID,
	addressID userent.AddressID,
) (address *userent.RegisteredAddress, err error) {
	return s.repo.GetUserAddress(ctx, userID, addressID)
}
