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

func TestAddUserAddressSuccess(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	address, err := userent.NewAddress(
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

	// Use case outputs
	expectedAddressID, err := uservo.NewAddressID(1)
	require.Nil(err)

	repo := &mocks.UserRepository{}
	repo.
		On("AddUserAddress", ctx, userID, address).
		Return(expectedAddressID, nil)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	addressID, err := useCases.AddUserAddress(ctx, userID, address)
	require.Nil(err)
	require.True(addressID.Equals(expectedAddressID))
}

func TestAddUserAddressFail(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	address, err := userent.NewAddress(
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

	// Use case outputs
	expectedErr := resperr.WithStatusCode(
		errors.New("test error"),
		http.StatusBadRequest,
	)

	repo := &mocks.UserRepository{}
	repo.
		On("AddUserAddress", ctx, userID, address).
		Return(
			uservo.AddressID(-1),
			expectedErr,
		)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	_, err = useCases.AddUserAddress(ctx, userID, address)
	require.ErrorIs(err, expectedErr)
}
