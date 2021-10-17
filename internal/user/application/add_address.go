package usersvc

import (
	"context"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (s UserService) AddAddress(
	ctx context.Context,
	userID uservo.UserID,
	address *userent.Address,
) (addressID uservo.AddressID, err error) {
	return s.repo.AddUserAddress(ctx, userID, address)
}
