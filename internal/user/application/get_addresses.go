package usersvc

import (
	"context"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (s UserService) GetAddresses(
	ctx context.Context,
	userID uservo.UserID,
	addressID uservo.AddressID,
) (addresses []*userent.RegisteredAddress, err error) {
	return s.repo.GetUserAddresses(ctx, userID)
}
