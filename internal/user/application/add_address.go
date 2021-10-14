package usersvc

import (
	"context"
	userent "ifoodish-store/internal/domain/entity"
)

func (s UserService) AddAddress(
	ctx context.Context,
	userID userent.UserID,
	address *userent.Address,
) (addressID userent.AddressID, err error) {
	return s.repo.AddUserAddress(ctx, userID, address)
}
