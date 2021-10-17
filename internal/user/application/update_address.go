package usersvc

import (
	"context"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (s UserService) UpdateAddress(
	ctx context.Context,
	userID uservo.UserID,
	address *userent.RegisteredAddress,
) (err error) {
	return s.repo.SaveUserAddress(ctx, userID, address)
}
