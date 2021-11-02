package useruc_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"ifoodish-store/mocks"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
	useruc "ifoodish-store/services/user/usecase"

	"github.com/carlmjohnson/resperr"
	"github.com/stretchr/testify/require"
)

func TestGetUserAddressesSuccess(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	// Use case outputs
	expectedAddress, err := userent.NewAddress(
		"Street ABCD",
		"Espirito Santo",
		"Jose dos Campos",
		"Rio de Janeiro",
		"Complement",
		"11111",
		"23970000",
		"-23.307577",
		"-44.754146",
	)
	require.Nil(err)

	expectedAddressID := uservo.GenerateNewAddressID()

	expectedRegisteredAddress, err := userent.NewRegisteredAddress(
		expectedAddressID.String(),
		expectedAddress,
	)
	require.Nil(err)

	repo := &mocks.UserRepository{}
	repo.
		On("GetUserAddresses", ctx, userID).
		Return([]userent.RegisteredAddress{expectedRegisteredAddress}, nil)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	addresses, err := useCases.GetUserAddresses(ctx, userID)
	require.Nil(err)
	require.Len(addresses, 1)
	require.EqualValues([]userent.RegisteredAddress{expectedRegisteredAddress}, addresses)
}

func TestGetUserAddressesFail(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	// Use case outputs
	expectedErr := resperr.WithStatusCode(
		errors.New("test error"),
		http.StatusBadRequest,
	)

	repo := &mocks.UserRepository{}
	repo.
		On("GetUserAddresses", ctx, userID).
		Return(
			[]userent.RegisteredAddress{},
			expectedErr,
		)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	addresses, err := useCases.GetUserAddresses(ctx, userID)
	require.ErrorIs(err, expectedErr)
	require.Len(addresses, 0)
}
