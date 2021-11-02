package useruc

import (
	"context"
	"fmt"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
)

func (s UserUseCases) GetUserAddresses(
	ctx context.Context,
	userID uservo.UserID,
) (addresses []userent.RegisteredAddress, err error) {
	userAddresses, err := s.repo.GetUserAddresses(ctx, userID)
	if err != nil {
		return userAddresses, fmt.Errorf("error getting user addresses: %w", err)
	}
	return userAddresses, nil
}
