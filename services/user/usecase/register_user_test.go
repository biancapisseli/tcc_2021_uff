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

func TestRegisterUserSuccess(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	user, err := userent.NewUser(
		"Lala Lele",
		"lala@lala.com",
		"5524543211234",
	)
	require.Nil(err)

	password, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	passwordConfirm, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	// services mocked values
	encodedPassword := uservo.NewPasswordEncoded("lalala")

	// Use case outputs
	expectedUserID := uservo.GenerateNewUserID()

	repo := &mocks.UserRepository{}
	repo.
		On("AddUser", ctx, user, encodedPassword).
		Return(expectedUserID, nil)

	encoder := &mocks.PasswordEncoder{}
	encoder.
		On("EncodePassword", password).
		Return(encodedPassword, nil)

	useCases := useruc.New(repo, encoder)

	userID, err := useCases.RegisterUser(ctx, user, password, passwordConfirm)
	require.Nil(err)
	require.True(userID.Equals(expectedUserID))
}

func TestRegisterUserPasswordNotEqualFail(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	user, err := userent.NewUser(
		"Lala Lele",
		"lala@lala.com",
		"5524543211234",
	)
	require.Nil(err)

	password, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	passwordConfirm, err := uservo.NewPasswordRaw("3213219")
	require.Nil(err)

	// Use case outputs

	repo := &mocks.UserRepository{}
	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	_, err = useCases.RegisterUser(ctx, user, password, passwordConfirm)
	require.Equal(http.StatusBadRequest, resperr.StatusCode(err))

}

func TestRegisterUserEncodePasswordFail(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	user, err := userent.NewUser(
		"Lala Lele",
		"lala@lala.com",
		"5524543211234",
	)
	require.Nil(err)

	password, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	passwordConfirm, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	// Use case outputs
	expectedErr := resperr.WithStatusCode(
		errors.New("test error"),
		http.StatusBadRequest,
	)

	repo := &mocks.UserRepository{}

	encoder := &mocks.PasswordEncoder{}
	encoder.
		On("EncodePassword", password).
		Return(uservo.PasswordEncoded(""), expectedErr)

	useCases := useruc.New(repo, encoder)

	_, err = useCases.RegisterUser(ctx, user, password, passwordConfirm)
	require.ErrorIs(err, expectedErr)

}

func TestRegisterUserAddUserFail(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	user, err := userent.NewUser(
		"Lala Lele",
		"lala@lala.com",
		"5524543211234",
	)
	require.Nil(err)

	password, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	passwordConfirm, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	// services mocked values
	encodedPassword := uservo.NewPasswordEncoded("lalala")

	// Use case outputs
	expectedErr := resperr.WithStatusCode(
		errors.New("test error"),
		http.StatusBadRequest,
	)

	repo := &mocks.UserRepository{}
	repo.
		On("AddUser", ctx, user, encodedPassword).
		Return(uservo.UserID{}, expectedErr)

	encoder := &mocks.PasswordEncoder{}
	encoder.
		On("EncodePassword", password).
		Return(encodedPassword, nil)

	useCases := useruc.New(repo, encoder)

	_, err = useCases.RegisterUser(ctx, user, password, passwordConfirm)
	require.ErrorIs(err, expectedErr)
}
