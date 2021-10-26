package usersvc

import (
	"context"
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (s UserService) AddUserAddress(
	ctx context.Context,
	userID uservo.UserID,
	address userent.Address,
) (addressID uservo.AddressID, err error) {
	addressID, err = s.repo.AddUserAddress(ctx, userID, address)
	if err != nil {
		return addressID, fmt.Errorf("error adding new user address: %w", err)
	}
	return addressID, nil
}
