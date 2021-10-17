package usersvc

import (
	"context"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (s UserService) GetAddress(
	ctx context.Context,
	userID uservo.UserID,
	addressID uservo.AddressID,
) (address *userent.RegisteredAddress, err error) {
	return s.repo.GetUserAddress(ctx, userID, addressID)
}
