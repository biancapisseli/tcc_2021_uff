package useruc_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	uservo "ifoodish-store/services/user/domain/valueobject"
	"ifoodish-store/services/user/mocks"
	useruc "ifoodish-store/services/user/usecase"

	"github.com/carlmjohnson/resperr"
	"github.com/stretchr/testify/require"
)

func TestRemoveUserAddressSuccess(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	addressID, err := uservo.NewAddressID(1)
	require.Nil(err)

	repo := &mocks.UserRepository{}
	repo.
		On("RemoveUserAddress", ctx, userID, addressID).
		Return(nil)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	err = useCases.RemoveUserAddress(ctx, userID, addressID)
	require.Nil(err)
}

func TestRemoveUserAddressFail(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	addressID, err := uservo.NewAddressID(1)
	require.Nil(err)

	// Use case outputs
	expectedErr := resperr.WithStatusCode(
		errors.New("test error"),
		http.StatusBadRequest,
	)

	repo := &mocks.UserRepository{}
	repo.
		On("RemoveUserAddress", ctx, userID, addressID).
		Return(expectedErr)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	err = useCases.RemoveUserAddress(ctx, userID, addressID)
	require.ErrorIs(err, expectedErr)
}
