package useruc

import (
	"context"
	"fmt"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (s UserUseCases) RemoveUserAddress(
	ctx context.Context,
	userID uservo.UserID,
	addressID uservo.AddressID,
) (err error) {
	err = s.repo.RemoveUserAddress(ctx, userID, addressID)
	if err != nil {
		return fmt.Errorf("error removing user address: %w", err)
	}
	return nil
}
