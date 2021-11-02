package useruc

import (
	"context"
	"fmt"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
)

func (s UserUseCases) AddUserAddress(
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
