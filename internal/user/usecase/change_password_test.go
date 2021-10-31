package useruc_test

import (
	"context"
	"errors"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/internal/user/mocks"
	useruc "ifoodish-store/internal/user/usecase"
	"net/http"
	"testing"

	"github.com/carlmjohnson/resperr"
	"github.com/stretchr/testify/require"
)

func TestChangePasswordSuccess(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	currentPassword, err := uservo.NewPasswordRaw("123123")
	require.Nil(err)

	newPassword, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	newPasswordConfirm, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	// Internal mocked values
	encodedCurrentPassword := uservo.NewPasswordEncoded("lalala")

	encodedNewPassword := uservo.NewPasswordEncoded("lelele")

	user, err := userent.NewUser("lalala", "lala@lala.com", "241234512345")
	require.Nil(err)

	registeredUser, err := userent.NewRegisteredUser(
		userID.String(),
		user,
	)
	require.Nil(err)

	repo := &mocks.UserRepository{}
	repo.
		On("GetUserInfo", ctx, userID).
		Return(registeredUser, nil)
	repo.
		On("GetUserByEmailAndPassword", ctx, user.Email, encodedCurrentPassword).
		Return(registeredUser, nil)
	repo.
		On("UpdatePassword", ctx, userID, encodedNewPassword).
		Return(nil)

	encoder := &mocks.PasswordEncoder{}
	encoder.
		On("EncodePassword", currentPassword).
		Return(encodedCurrentPassword, nil)
	encoder.
		On("EncodePassword", newPassword).
		Return(encodedNewPassword, nil)

	useCases := useruc.New(repo, encoder)

	err = useCases.ChangePassword(
		ctx,
		userID,
		currentPassword,
		newPassword,
		newPasswordConfirm,
	)
	require.Nil(err)
}

func TestChangePasswordNotEqualFail(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	currentPassword, err := uservo.NewPasswordRaw("123123")
	require.Nil(err)

	newPassword, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	newPasswordConfirm, err := uservo.NewPasswordRaw("3213219")
	require.Nil(err)

	repo := &mocks.UserRepository{}
	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	err = useCases.ChangePassword(
		ctx,
		userID,
		currentPassword,
		newPassword,
		newPasswordConfirm,
	)
	require.Equal(http.StatusBadRequest, resperr.StatusCode(err))
}

func TestChangePasswordGetUserInfoFail(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	currentPassword, err := uservo.NewPasswordRaw("123123")
	require.Nil(err)

	newPassword, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	newPasswordConfirm, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	// Use case outputs
	expectedErr := resperr.WithStatusCode(
		errors.New("test error"),
		http.StatusBadRequest,
	)

	repo := &mocks.UserRepository{}
	repo.
		On("GetUserInfo", ctx, userID).
		Return(userent.RegisteredUser{}, expectedErr)

	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	err = useCases.ChangePassword(
		ctx,
		userID,
		currentPassword,
		newPassword,
		newPasswordConfirm,
	)
	require.ErrorIs(err, expectedErr)
}

func TestChangePasswordEncodeCurrentPasswordFail(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	currentPassword, err := uservo.NewPasswordRaw("123123")
	require.Nil(err)

	newPassword, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	newPasswordConfirm, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	// Internal mocked values
	user, err := userent.NewUser("lalala", "lala@lala.com", "241234512345")
	require.Nil(err)

	registeredUser, err := userent.NewRegisteredUser(
		userID.String(),
		user,
	)
	require.Nil(err)

	// Use case outputs
	expectedErr := resperr.WithStatusCode(
		errors.New("test error"),
		http.StatusBadRequest,
	)

	repo := &mocks.UserRepository{}
	repo.
		On("GetUserInfo", ctx, userID).
		Return(registeredUser, nil)

	encoder := &mocks.PasswordEncoder{}
	encoder.
		On("EncodePassword", currentPassword).
		Return(uservo.PasswordEncoded(""), expectedErr)

	useCases := useruc.New(repo, encoder)

	err = useCases.ChangePassword(
		ctx,
		userID,
		currentPassword,
		newPassword,
		newPasswordConfirm,
	)
	require.ErrorIs(err, expectedErr)
}

func TestChangePasswordAuthFail(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	currentPassword, err := uservo.NewPasswordRaw("123123")
	require.Nil(err)

	newPassword, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	newPasswordConfirm, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	// Internal mocked values
	encodedCurrentPassword := uservo.NewPasswordEncoded("lalala")

	user, err := userent.NewUser("lalala", "lala@lala.com", "241234512345")
	require.Nil(err)

	registeredUser, err := userent.NewRegisteredUser(
		userID.String(),
		user,
	)
	require.Nil(err)

	// Use case outputs
	expectedErr := resperr.WithStatusCode(
		errors.New("test error"),
		http.StatusBadRequest,
	)

	repo := &mocks.UserRepository{}
	repo.
		On("GetUserInfo", ctx, userID).
		Return(registeredUser, nil)
	repo.
		On("GetUserByEmailAndPassword", ctx, user.Email, encodedCurrentPassword).
		Return(userent.RegisteredUser{}, expectedErr)

	encoder := &mocks.PasswordEncoder{}
	encoder.
		On("EncodePassword", currentPassword).
		Return(encodedCurrentPassword, nil)

	useCases := useruc.New(repo, encoder)

	err = useCases.ChangePassword(
		ctx,
		userID,
		currentPassword,
		newPassword,
		newPasswordConfirm,
	)
	require.ErrorIs(err, expectedErr)
}

func TestChangePasswordEncodeNewPasswordFail(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	currentPassword, err := uservo.NewPasswordRaw("123123")
	require.Nil(err)

	newPassword, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	newPasswordConfirm, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	// Internal mocked values
	encodedCurrentPassword := uservo.NewPasswordEncoded("lalala")

	user, err := userent.NewUser("lalala", "lala@lala.com", "241234512345")
	require.Nil(err)

	registeredUser, err := userent.NewRegisteredUser(
		userID.String(),
		user,
	)
	require.Nil(err)

	// Use case outputs
	expectedErr := resperr.WithStatusCode(
		errors.New("test error"),
		http.StatusBadRequest,
	)

	repo := &mocks.UserRepository{}
	repo.
		On("GetUserInfo", ctx, userID).
		Return(registeredUser, nil)
	repo.
		On("GetUserByEmailAndPassword", ctx, user.Email, encodedCurrentPassword).
		Return(registeredUser, nil)

	encoder := &mocks.PasswordEncoder{}
	encoder.
		On("EncodePassword", currentPassword).
		Return(encodedCurrentPassword, nil)
	encoder.
		On("EncodePassword", newPassword).
		Return(uservo.PasswordEncoded(""), expectedErr)

	useCases := useruc.New(repo, encoder)

	err = useCases.ChangePassword(
		ctx,
		userID,
		currentPassword,
		newPassword,
		newPasswordConfirm,
	)
	require.ErrorIs(err, expectedErr)
}

func TestChangePasswordUpdatePasswordFail(t *testing.T) {
	require := require.New(t)

	// Use case inputs
	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	currentPassword, err := uservo.NewPasswordRaw("123123")
	require.Nil(err)

	newPassword, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	newPasswordConfirm, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	// Internal mocked values
	encodedCurrentPassword := uservo.NewPasswordEncoded("lalala")

	encodedNewPassword := uservo.NewPasswordEncoded("lelele")

	user, err := userent.NewUser("lalala", "lala@lala.com", "241234512345")
	require.Nil(err)

	registeredUser, err := userent.NewRegisteredUser(
		userID.String(),
		user,
	)
	require.Nil(err)

	// Use case outputs
	expectedErr := resperr.WithStatusCode(
		errors.New("test error"),
		http.StatusBadRequest,
	)

	repo := &mocks.UserRepository{}
	repo.
		On("GetUserInfo", ctx, userID).
		Return(registeredUser, nil)
	repo.
		On("GetUserByEmailAndPassword", ctx, user.Email, encodedCurrentPassword).
		Return(registeredUser, nil)
	repo.
		On("UpdatePassword", ctx, userID, encodedNewPassword).
		Return(expectedErr)

	encoder := &mocks.PasswordEncoder{}
	encoder.
		On("EncodePassword", currentPassword).
		Return(encodedCurrentPassword, nil)
	encoder.
		On("EncodePassword", newPassword).
		Return(encodedNewPassword, nil)

	useCases := useruc.New(repo, encoder)

	err = useCases.ChangePassword(
		ctx,
		userID,
		currentPassword,
		newPassword,
		newPasswordConfirm,
	)
	require.ErrorIs(err, expectedErr)
}
