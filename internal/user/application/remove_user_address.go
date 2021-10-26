package usersvc

import (
	"context"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (s UserService) RemoveUserAddress(
	ctx context.Context,
	userID uservo.UserID,
	addressID uservo.AddressID,
) (err error) {
	return s.repo.RemoveUserAddress(ctx, userID, addressID)
}
