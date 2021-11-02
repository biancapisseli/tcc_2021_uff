package userhttpcontroller_test

import (
	"context"
	"encoding/json"
	"errors"
	uservo "ifoodish-store/services/user/domain/valueobject"
	userhttpcontroller "ifoodish-store/services/user/infrastructure/controller/http"
	"ifoodish-store/services/user/mocks"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestChangePasswordSuccess(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	currentPassword, err := uservo.NewPasswordRaw("123123")
	require.Nil(err)

	newPassword, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	newPasswordConfirm, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	body := userhttpcontroller.ChangePasswordBody{
		CurrentPassword:    currentPassword,
		NewPassword:        newPassword,
		NewPasswordConfirm: newPasswordConfirm,
	}

	useCases := &mocks.UserUseCases{}
	useCases.On("ChangePassword",
		ctx,
		userID,
		currentPassword,
		newPassword,
		newPasswordConfirm,
	).Return(nil)

	req := &mocks.Request{}
	req.On("ParseBodyParams",
		mock.AnythingOfType("*userhttpcontroller.ChangePasswordBody"),
	).Return(nil).Run(func(args mock.Arguments) {
		argBody := args.Get(0).(*userhttpcontroller.ChangePasswordBody)
		*argBody = body
	})
	req.On("Context").Return(ctx)
	req.On("GetUserID").Return(userID, nil)

	controller := userhttpcontroller.New(useCases)

	err = controller.ChangePassword(req)
	require.Nil(err)
}

func TestChangePasswordUserIDError(t *testing.T) {
	require := require.New(t)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}

	req := &mocks.Request{}
	req.On("GetUserID").Return(uservo.GenerateNewUserID(), expectedErr)

	controller := userhttpcontroller.New(useCases)

	err := controller.ChangePassword(req)
	require.ErrorIs(err, expectedErr)
}

func TestChangePasswordParseBodyFail(t *testing.T) {
	require := require.New(t)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}

	req := &mocks.Request{}
	req.On("GetUserID").Return(uservo.GenerateNewUserID(), nil)
	req.On("ParseBodyParams",
		mock.AnythingOfType("*userhttpcontroller.ChangePasswordBody"),
	).Return(expectedErr)

	controller := userhttpcontroller.New(useCases)

	err := controller.ChangePassword(req)
	require.ErrorIs(err, expectedErr)
}

func TestChangePasswordUseCaseFail(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	currentPassword, err := uservo.NewPasswordRaw("123123")
	require.Nil(err)

	newPassword, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	newPasswordConfirm, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	body := userhttpcontroller.ChangePasswordBody{
		CurrentPassword:    currentPassword,
		NewPassword:        newPassword,
		NewPasswordConfirm: newPasswordConfirm,
	}

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}
	useCases.On("ChangePassword",
		ctx,
		userID,
		currentPassword,
		newPassword,
		newPasswordConfirm,
	).Return(expectedErr)

	req := &mocks.Request{}
	req.On("ParseBodyParams",
		mock.AnythingOfType("*userhttpcontroller.ChangePasswordBody"),
	).Return(nil).Run(func(args mock.Arguments) {
		argBody := args.Get(0).(*userhttpcontroller.ChangePasswordBody)
		*argBody = body
	})
	req.On("Context").Return(ctx)
	req.On("GetUserID").Return(userID, nil)

	controller := userhttpcontroller.New(useCases)

	err = controller.ChangePassword(req)
	require.ErrorIs(err, expectedErr)
}

func TestChangePasswordUnmarshalDomainFail(t *testing.T) {
	require := require.New(t)

	for index, tc := range []struct {
		current    string
		new        string
		newConfirm string
		err        error
	}{{
		current:    strings.Repeat("1", uservo.MaxRawPasswordLength),
		new:        strings.Repeat("1", uservo.MaxRawPasswordLength),
		newConfirm: strings.Repeat("1", uservo.MaxRawPasswordLength),
		err:        nil,
	}, {
		current:    strings.Repeat("1", uservo.MaxRawPasswordLength+1),
		new:        strings.Repeat("1", uservo.MaxRawPasswordLength),
		newConfirm: strings.Repeat("1", uservo.MaxRawPasswordLength),
		err:        uservo.ErrRawPasswordMaxLength,
	}, {
		current:    strings.Repeat("1", uservo.MaxRawPasswordLength),
		new:        strings.Repeat("1", uservo.MaxRawPasswordLength+1),
		newConfirm: strings.Repeat("1", uservo.MaxRawPasswordLength),
		err:        uservo.ErrRawPasswordMaxLength,
	}, {
		current:    strings.Repeat("1", uservo.MaxRawPasswordLength),
		new:        strings.Repeat("1", uservo.MaxRawPasswordLength),
		newConfirm: strings.Repeat("1", uservo.MaxRawPasswordLength+1),
		err:        uservo.ErrRawPasswordMaxLength,
	}} {
		b, err := json.Marshal(userhttpcontroller.ChangePasswordBody{
			CurrentPassword:    uservo.PasswordRaw(tc.current),
			NewPassword:        uservo.PasswordRaw(tc.new),
			NewPasswordConfirm: uservo.PasswordRaw(tc.newConfirm),
		})
		require.Nil(err)

		var body userhttpcontroller.ChangePasswordBody
		err = json.Unmarshal(b, &body)
		require.ErrorIs(err, tc.err, "index %d", index)
	}
}

func TestChangePasswordUnmarshalFail(t *testing.T) {
	require := require.New(t)

	for index, tc := range []struct {
		data string
		err  error
	}{{
		data: `{
			"current_password":"123123",
			"new_password":"321321",
			"new_password_confirm":"321321",
		}`,
		err: &json.SyntaxError{},
	}, {
		data: `{
			"current_password":123123,
			"new_password":"321321",
			"new_password_confirm":"321321"
		}`,
		err: &json.UnmarshalTypeError{},
	}} {
		var body userhttpcontroller.ChangePasswordBody
		err := json.Unmarshal([]byte(tc.data), &body)
		require.ErrorAs(err, &tc.err, "index %d", index)
	}
}
