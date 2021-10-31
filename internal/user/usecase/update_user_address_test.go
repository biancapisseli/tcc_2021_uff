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

func TestUpdateUserAddressSuccess(t *testing.T) {
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

	registeredAddress, err := userent.NewRegisteredAddress(1, address)
	require.Nil(err)

	repo := &mocks.UserRepository{}
	repo.
		On("SaveUserAddress", ctx, userID, registeredAddress).
		Return(nil)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	err = useCases.UpdateUserAddress(ctx, userID, registeredAddress)
	require.Nil(err)
}

func TestUpdateUserAddressFail(t *testing.T) {
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

	registeredAddress, err := userent.NewRegisteredAddress(1, address)
	require.Nil(err)

	// Use case outputs
	expectedErr := resperr.WithStatusCode(
		errors.New("test error"),
		http.StatusBadRequest,
	)

	repo := &mocks.UserRepository{}
	repo.
		On("SaveUserAddress", ctx, userID, registeredAddress).
		Return(expectedErr)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	err = useCases.UpdateUserAddress(ctx, userID, registeredAddress)
	require.ErrorIs(err, expectedErr)
}
