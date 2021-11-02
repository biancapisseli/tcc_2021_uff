package userhttpcontroller_test

import (
	"context"
	"errors"
	"net/http"

	uservo "ifoodish-store/services/user/domain/valueobject"
	userhttpcontroller "ifoodish-store/services/user/infrastructure/controller/http"
	"ifoodish-store/services/user/mocks"
	"testing"

	"github.com/carlmjohnson/resperr"
	"github.com/stretchr/testify/require"
)

func TestRemoveUserAddressSuccess(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	addressID := uservo.GenerateNewAddressID()

	useCases := &mocks.UserUseCases{}
	useCases.On("RemoveUserAddress", ctx, userID, addressID).Return(nil)

	req := &mocks.Request{}
	req.
		On("GetURLParam", "address_id").
		Return(addressID.String())
	req.On("Context").Return(ctx)
	req.On("GetUserID").Return(userID, nil)

	controller := userhttpcontroller.New(useCases)

	err := controller.RemoveUserAddress(req)
	require.Nil(err)
}

func TestRemoveUserAddressUserIDError(t *testing.T) {
	require := require.New(t)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}

	req := &mocks.Request{}
	req.On("GetUserID").Return(uservo.GenerateNewUserID(), expectedErr)

	controller := userhttpcontroller.New(useCases)

	err := controller.RemoveUserAddress(req)
	require.ErrorIs(err, expectedErr)
}

func TestRemoveUserAddressParamFail(t *testing.T) {
	require := require.New(t)

	useCases := &mocks.UserUseCases{}

	req := &mocks.Request{}
	req.On("GetUserID").Return(uservo.GenerateNewUserID(), nil)
	req.
		On("GetURLParam", "address_id").
		Return("")

	controller := userhttpcontroller.New(useCases)

	err := controller.RemoveUserAddress(req)
	require.Equal(http.StatusBadRequest, resperr.StatusCode(err))
}

func TestRemoveUserAddressUseCaseFail(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	addressID := uservo.GenerateNewAddressID()

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}
	useCases.On("RemoveUserAddress", ctx, userID, addressID).Return(expectedErr)

	req := &mocks.Request{}
	req.
		On("GetURLParam", "address_id").
		Return(addressID.String())
	req.On("Context").Return(ctx)
	req.On("GetUserID").Return(userID, nil)

	controller := userhttpcontroller.New(useCases)

	err := controller.RemoveUserAddress(req)
	require.ErrorIs(err, expectedErr)
}
