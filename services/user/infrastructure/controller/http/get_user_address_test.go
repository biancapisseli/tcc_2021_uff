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

func TestGetUserAddressSuccess(t *testing.T) {
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
	require.Nil(err)

	expectedRegAddress, err := userent.NewRegisteredAddress(addressID.String(), address)
	require.Nil(err)

	useCases := &mocks.UserUseCases{}
	useCases.On("GetUserAddress", ctx, userID, addressID).Return(expectedRegAddress, nil)

	req := &mocks.Request{}
	req.
		On("GetURLParam", "address_id").
		Return(addressID.String())
	req.On("Context").Return(ctx)
	req.On("GetUserID").Return(userID, nil)

	controller := userhttpcontroller.New(useCases)

	regAddress, err := controller.GetUserAddress(req)
	require.Nil(err)
	require.True(regAddress.ID.Equals(addressID))
}

func TestGetUserAddressUserIDError(t *testing.T) {
	require := require.New(t)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}

	req := &mocks.Request{}
	req.On("GetUserID").Return(uservo.GenerateNewUserID(), expectedErr)

	controller := userhttpcontroller.New(useCases)

	_, err := controller.GetUserAddress(req)
	require.ErrorIs(err, expectedErr)
}

// func TestGetUserAddressParseParamsFail(t *testing.T) {
// 	require := require.New(t)

// 	expectedErr := errors.New("test error")

// 	useCases := &mocks.UserUseCases{}

// 	req := &mocks.Request{}
// 	req.On("GetUserID").Return(uservo.GenerateNewUserID(), nil)
// 	req.On("ParseBodyParams",
// 		mock.AnythingOfType("*userent.Address"),
// 	).Return(expectedErr)

// 	controller := userhttpcontroller.New(useCases)

// 	_, err := controller.GetUserAddress(req)
// 	require.ErrorIs(err, expectedErr)
// }

// func TestGetUserAddressUseCaseFail(t *testing.T) {
// 	require := require.New(t)

// 	ctx := context.Background()

// 	userID := uservo.GenerateNewUserID()

// 	address, err := userent.NewAddress(
// 		"Street ABCD",
// 		"Espirito Santo",
// 		"Jose dos Campos",
// 		"Rio de Janeiro",
// 		"Complement",
// 		"11111",
// 		"23970000",
// 		"-23.307577",
// 		"-44.754146",
// 	)
// 	require.Nil(err)

// 	expectedErr := errors.New("test error")

// 	useCases := &mocks.UserUseCases{}
// 	useCases.On("GetUserAddress", ctx, userID, address).Return(uservo.AddressID(-1), expectedErr)

// 	req := &mocks.Request{}
// 	req.On("ParseBodyParams",
// 		mock.AnythingOfType("*userent.Address"),
// 	).Return(nil).Run(func(args mock.Arguments) {
// 		argAddress := args.Get(0).(*userent.Address)
// 		*argAddress = address
// 	})
// 	req.On("Context").Return(ctx)
// 	req.On("GetUserID").Return(userID, nil)

// 	controller := userhttpcontroller.New(useCases)

// 	_, err = controller.GetUserAddress(req)
// 	require.ErrorIs(err, expectedErr)
// }
