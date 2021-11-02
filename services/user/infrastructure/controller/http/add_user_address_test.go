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

func TestAddUserAddressSuccess(t *testing.T) {
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

	expectedAddressID := uservo.GenerateNewAddressID()

	useCases := &mocks.UserUseCases{}
	useCases.On("AddUserAddress", ctx, userID, address).Return(expectedAddressID, nil)

	req := &mocks.Request{}
	req.On("ParseBody",
		mock.AnythingOfType("*userent.Address"),
	).Return(nil).Run(func(args mock.Arguments) {
		argAddress := args.Get(0).(*userent.Address)
		*argAddress = address
	})
	req.On("Context").Return(ctx)
	req.On("GetUserID").Return(userID, nil)

	controller := userhttpcontroller.New(useCases)

	addressID, err := controller.AddUserAddress(req)
	require.Nil(err)
	require.True(addressID.Equals(expectedAddressID))
}

func TestAddUserAddressUserIDError(t *testing.T) {
	require := require.New(t)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}

	req := &mocks.Request{}
	req.On("GetUserID").Return(uservo.GenerateNewUserID(), expectedErr)

	controller := userhttpcontroller.New(useCases)

	_, err := controller.AddUserAddress(req)
	require.ErrorIs(err, expectedErr)
}

func TestAddUserAddressParseBodyFail(t *testing.T) {
	require := require.New(t)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}

	req := &mocks.Request{}
	req.On("GetUserID").Return(uservo.GenerateNewUserID(), nil)
	req.On("ParseBody",
		mock.AnythingOfType("*userent.Address"),
	).Return(expectedErr)

	controller := userhttpcontroller.New(useCases)

	_, err := controller.AddUserAddress(req)
	require.ErrorIs(err, expectedErr)
}

func TestAddUserAddressUseCaseFail(t *testing.T) {
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

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}
	useCases.
		On("AddUserAddress", ctx, userID, address).
		Return(addressID, expectedErr)

	req := &mocks.Request{}
	req.On("ParseBody",
		mock.AnythingOfType("*userent.Address"),
	).Return(nil).Run(func(args mock.Arguments) {
		argAddress := args.Get(0).(*userent.Address)
		*argAddress = address
	})
	req.On("Context").Return(ctx)
	req.On("GetUserID").Return(userID, nil)

	controller := userhttpcontroller.New(useCases)

	_, err = controller.AddUserAddress(req)
	require.ErrorIs(err, expectedErr)
}
