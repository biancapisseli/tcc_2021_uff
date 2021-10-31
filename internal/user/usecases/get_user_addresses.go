package useruc

import (
	"context"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (s UserUseCases) GetUserAddresses(
	ctx context.Context,
	userID uservo.UserID,
) (addresses []userent.RegisteredAddress, err error) {
	return s.repo.GetUserAddresses(ctx, userID)
}