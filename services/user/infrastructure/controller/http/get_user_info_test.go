package userhttpcontroller_test

import (
	"context"
	"errors"
	userent "ifoodish-store/services/user/domain/entity"

	"ifoodish-store/mocks"
	uservo "ifoodish-store/services/user/domain/valueobject"
	userhttpcontroller "ifoodish-store/services/user/infrastructure/controller/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUserInfoSuccess(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	user, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	expectedRegUser, err := userent.NewRegisteredUser(userID.String(), user)
	require.Nil(err)

	useCases := &mocks.UserUseCases{}
	useCases.On("GetUserInfo", ctx, userID).Return(expectedRegUser, nil)

	req := &mocks.Request{}
	req.On("Context").Return(ctx)
	req.On("GetUserID").Return(userID, nil)

	controller := userhttpcontroller.New(useCases)

	regAddress, err := controller.GetUserInfo(req)
	require.Nil(err)
	require.True(regAddress.ID.Equals(userID))
}

func TestGetUserInfoUserIDError(t *testing.T) {
	require := require.New(t)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}

	req := &mocks.Request{}
	req.On("GetUserID").Return(uservo.GenerateNewUserID(), expectedErr)

	controller := userhttpcontroller.New(useCases)

	_, err := controller.GetUserInfo(req)
	require.ErrorIs(err, expectedErr)
}

func TestGetUserInfoUseCaseFail(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	user, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	expectedRegUser, err := userent.NewRegisteredUser(userID.String(), user)
	require.Nil(err)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}
	useCases.On("GetUserInfo", ctx, userID).Return(expectedRegUser, expectedErr)

	req := &mocks.Request{}
	req.On("Context").Return(ctx)
	req.On("GetUserID").Return(userID, nil)

	controller := userhttpcontroller.New(useCases)

	_, err = controller.GetUserInfo(req)
	require.ErrorIs(err, expectedErr)
}
