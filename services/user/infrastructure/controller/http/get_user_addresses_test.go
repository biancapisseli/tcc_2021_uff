package userhttpcontroller_test

import (
	"context"
	"errors"
	userent "ifoodish-store/services/user/domain/entity"

	uservo "ifoodish-store/services/user/domain/valueobject"
	userhttpcontroller "ifoodish-store/services/user/infrastructure/controller/http"
	"ifoodish-store/services/user/mocks"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUserAddressesSuccess(t *testing.T) {
	require := require.New(t)

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

	addressID := uservo.GenerateNewAddressID()

	expectedRegAddress, err := userent.NewRegisteredAddress(addressID.String(), address)
	require.Nil(err)

	useCases := &mocks.UserUseCases{}
	useCases.
		On("GetUserAddresses", ctx, userID).
		Return([]userent.RegisteredAddress{expectedRegAddress}, nil)

	req := &mocks.Request{}
	req.On("Context").Return(ctx)
	req.On("GetUserID").Return(userID, nil)

	controller := userhttpcontroller.New(useCases)

	regAddresses, err := controller.GetUserAddresses(req)
	require.Nil(err)
	require.Len(regAddresses, 1)
	require.True(regAddresses[0].ID.Equals(addressID))
}

func TestGetUserAddressesUserIDError(t *testing.T) {
	require := require.New(t)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}

	req := &mocks.Request{}
	req.On("GetUserID").Return(uservo.GenerateNewUserID(), expectedErr)

	controller := userhttpcontroller.New(useCases)

	_, err := controller.GetUserAddresses(req)
	require.ErrorIs(err, expectedErr)
}

func TestGetUserAddressesUseCaseFail(t *testing.T) {
	require := require.New(t)

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

	addressID := uservo.GenerateNewAddressID()

	expectedRegAddress, err := userent.NewRegisteredAddress(addressID.String(), address)
	require.Nil(err)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}
	useCases.
		On("GetUserAddresses", ctx, userID).
		Return([]userent.RegisteredAddress{expectedRegAddress}, expectedErr)

	req := &mocks.Request{}
	req.On("Context").Return(ctx)
	req.On("GetUserID").Return(userID, nil)

	controller := userhttpcontroller.New(useCases)

	_, err = controller.GetUserAddresses(req)
	require.ErrorIs(err, expectedErr)
}
