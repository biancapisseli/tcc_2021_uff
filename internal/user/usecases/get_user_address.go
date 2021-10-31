package useruc

import (
	"context"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (s UserUseCases) GetUserAddress(
	ctx context.Context,
	userID uservo.UserID,
	addressID uservo.AddressID,
) (address userent.RegisteredAddress, err error) {
	return s.repo.GetUserAddress(ctx, userID, addressID)
}
