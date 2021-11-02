package useruc

import (
	"context"
	"fmt"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
)

func (s UserUseCases) UpdateUserAddress(
	ctx context.Context,
	userID uservo.UserID,
	address userent.RegisteredAddress,
) (err error) {
	err = s.repo.SaveUserAddress(ctx, userID, address)
	if err != nil {
		return fmt.Errorf("error updating user address: %w", err)
	}
	return nil
}
