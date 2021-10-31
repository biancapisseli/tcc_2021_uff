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

func TestUpdateUserInfoSuccess(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	user, err := userent.NewUser(
		"Lala Lele",
		"lala@lala.com",
		"5524543211234",
	)
	require.Nil(err)

	registeredUser, err := userent.NewRegisteredUser(userID.String(), user)
	require.Nil(err)

	repo := &mocks.UserRepository{}
	repo.
		On("SaveUser", ctx, registeredUser).
		Return(nil)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	err = useCases.UpdateUserInfo(ctx, registeredUser)
	require.Nil(err)
}

func TestUpdateUserInfoFail(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	user, err := userent.NewUser(
		"Lala Lele",
		"lala@lala.com",
		"5524543211234",
	)
	require.Nil(err)

	registeredUser, err := userent.NewRegisteredUser(userID.String(), user)
	require.Nil(err)

	// Use case outputs
	expectedErr := resperr.WithStatusCode(
		errors.New("test error"),
		http.StatusBadRequest,
	)

	repo := &mocks.UserRepository{}
	repo.
		On("SaveUser", ctx, registeredUser).
		Return(expectedErr)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	err = useCases.UpdateUserInfo(ctx, registeredUser)
	require.ErrorIs(err, expectedErr)
}
