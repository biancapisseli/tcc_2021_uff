package usersvc

import (
	"context"
	userent "ifoodish-store/internal/domain/entity"
)

func (s UserService) UpdateAddress(
	ctx context.Context,
	userID userent.UserID,
	address *userent.RegisteredAddress,
) (err error) {
	return s.repo.SaveUserAddress(ctx, userID, address)
}
