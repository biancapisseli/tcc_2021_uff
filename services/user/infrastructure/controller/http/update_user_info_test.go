package userhttpcontroller_test

import (
	"context"
	"errors"
	"ifoodish-store/mocks"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
	userhttpcontroller "ifoodish-store/services/user/infrastructure/controller/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdateUserInfoSuccess(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	body, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	useCases := &mocks.UserUseCases{}
	useCases.On("UpdateUserInfo", ctx, userID, body).Return(nil)

	req := &mocks.Request{}
	req.On("ParseBody",
		mock.AnythingOfType("*userent.User"),
	).Return(nil).Run(func(args mock.Arguments) {
		argAddress := args.Get(0).(*userent.User)
		*argAddress = body
	})
	req.On("Context").Return(ctx)
	req.On("GetUserID").Return(userID, nil)

	controller := userhttpcontroller.New(useCases)

	err = controller.UpdateUserInfo(req)
	require.Nil(err)
}

func TestUpdateUserInfoUserIDError(t *testing.T) {
	require := require.New(t)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}

	req := &mocks.Request{}
	req.On("GetUserID").Return(uservo.GenerateNewUserID(), expectedErr)

	controller := userhttpcontroller.New(useCases)

	err := controller.UpdateUserInfo(req)
	require.ErrorIs(err, expectedErr)
}

func TestUpdateUserInfoParseBodyFail(t *testing.T) {
	require := require.New(t)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}

	req := &mocks.Request{}
	req.On("GetUserID").Return(uservo.GenerateNewUserID(), nil)
	req.On("ParseBody",
		mock.AnythingOfType("*userent.User"),
	).Return(expectedErr)

	controller := userhttpcontroller.New(useCases)

	err := controller.UpdateUserInfo(req)
	require.ErrorIs(err, expectedErr)
}

func TestUpdateUserInfoUseCaseFail(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	body, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}
	useCases.
		On("UpdateUserInfo", ctx, userID, body).
		Return(expectedErr)

	req := &mocks.Request{}
	req.On("ParseBody",
		mock.AnythingOfType("*userent.User"),
	).Return(nil).Run(func(args mock.Arguments) {
		argAddress := args.Get(0).(*userent.User)
		*argAddress = body
	})
	req.On("Context").Return(ctx)
	req.On("GetUserID").Return(userID, nil)

	controller := userhttpcontroller.New(useCases)

	err = controller.UpdateUserInfo(req)
	require.ErrorIs(err, expectedErr)
}
