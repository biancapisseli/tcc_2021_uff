package usersvc

import (
	"context"
	userent "ifoodish-store/internal/domain/entity"
)

func (s UserService) RemoveAddress(
	ctx context.Context,
	userID userent.UserID,
	addressID userent.AddressID,
) (err error) {
	return s.repo.RemoveUserAddress(ctx, userID, addressID)
}
