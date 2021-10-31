package useruc

import (
	"context"
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (s UserUseCases) GetUserAddress(
	ctx context.Context,
	userID uservo.UserID,
	addressID uservo.AddressID,
) (address userent.RegisteredAddress, err error) {
	userAddress, err := s.repo.GetUserAddress(ctx, userID, addressID)
	if err != nil {
		return userAddress, fmt.Errorf("error getting user address: %w", err)
	}
	return userAddress, nil
}
