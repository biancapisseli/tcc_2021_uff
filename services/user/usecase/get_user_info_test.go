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

func TestGetUserInfoSuccess(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	// Use case outputs
	expectedUser, err := userent.NewUser(
		"Lala Lele",
		"lala@lala.com",
		"5524543211234",
	)
	require.Nil(err)

	expectedRegisteredUser, err := userent.NewRegisteredUser(userID.String(), expectedUser)
	require.Nil(err)

	repo := &mocks.UserRepository{}
	repo.
		On("GetUserInfo", ctx, userID).
		Return(expectedRegisteredUser, nil)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	user, err := useCases.GetUserInfo(ctx, userID)
	require.Nil(err)
	require.EqualValues(expectedRegisteredUser, user)
}

func TestGetUserInfoFail(t *testing.T) {
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
		On("GetUserInfo", ctx, userID).
		Return(
			userent.RegisteredUser{},
			expectedErr,
		)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	_, err := useCases.GetUserInfo(ctx, userID)
	require.ErrorIs(err, expectedErr)
}
