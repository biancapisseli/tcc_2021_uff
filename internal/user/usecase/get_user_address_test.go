package useruc_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/internal/user/mocks"
	useruc "ifoodish-store/internal/user/usecase"

	"github.com/carlmjohnson/resperr"
	"github.com/stretchr/testify/require"
)

func TestGetUserAddressSuccess(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	addressID, err := uservo.NewAddressID(1)
	require.Nil(err)

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

	expectedRegisteredAddress, err := userent.NewRegisteredAddress(1, expectedAddress)
	require.Nil(err)

	repo := &mocks.UserRepository{}
	repo.
		On("GetUserAddress", ctx, userID, addressID).
		Return(expectedRegisteredAddress, nil)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	address, err := useCases.GetUserAddress(ctx, userID, addressID)
	require.Nil(err)
	require.EqualValues(expectedRegisteredAddress, address)
}

func TestGetUserAddressFail(t *testing.T) {
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
		On("GetUserAddress", ctx, userID, addressID).
		Return(
			userent.RegisteredAddress{},
			expectedErr,
		)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	_, err = useCases.GetUserAddress(ctx, userID, addressID)
	require.ErrorIs(err, expectedErr)
}
