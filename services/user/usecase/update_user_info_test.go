package useruc_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
	"ifoodish-store/services/user/mocks"
	useruc "ifoodish-store/services/user/usecase"

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

	repo := &mocks.UserRepository{}
	repo.
		On("SaveUser", ctx, userID, user).
		Return(nil)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	err = useCases.UpdateUserInfo(ctx, userID, user)
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

	// Use case outputs
	expectedErr := resperr.WithStatusCode(
		errors.New("test error"),
		http.StatusBadRequest,
	)

	repo := &mocks.UserRepository{}
	repo.
		On("SaveUser", ctx, userID, user).
		Return(expectedErr)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	err = useCases.UpdateUserInfo(ctx, userID, user)
	require.ErrorIs(err, expectedErr)
}
