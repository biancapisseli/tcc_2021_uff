package userhttpcontroller_test

import (
	"context"
	"errors"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
	userhttpcontroller "ifoodish-store/services/user/infrastructure/controller/http"
	"ifoodish-store/services/user/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdateUserAddressSuccess(t *testing.T) {
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

	body, err := userent.NewRegisteredAddress(addressID.String(), address)
	require.Nil(err)

	useCases := &mocks.UserUseCases{}
	useCases.On("UpdateUserAddress", ctx, userID, body).Return(nil)

	req := &mocks.Request{}
	req.On("ParseBody",
		mock.AnythingOfType("*userent.RegisteredAddress"),
	).Return(nil).Run(func(args mock.Arguments) {
		argAddress := args.Get(0).(*userent.RegisteredAddress)
		*argAddress = body
	})
	req.On("Context").Return(ctx)
	req.On("GetUserID").Return(userID, nil)

	controller := userhttpcontroller.New(useCases)

	err = controller.UpdateUserAddress(req)
	require.Nil(err)
}

func TestUpdateUserAddressUserIDError(t *testing.T) {
	require := require.New(t)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}

	req := &mocks.Request{}
	req.On("GetUserID").Return(uservo.GenerateNewUserID(), expectedErr)

	controller := userhttpcontroller.New(useCases)

	err := controller.UpdateUserAddress(req)
	require.ErrorIs(err, expectedErr)
}

func TestUpdateUserAddressParseBodyFail(t *testing.T) {
	require := require.New(t)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}

	req := &mocks.Request{}
	req.On("GetUserID").Return(uservo.GenerateNewUserID(), nil)
	req.On("ParseBody",
		mock.AnythingOfType("*userent.RegisteredAddress"),
	).Return(expectedErr)

	controller := userhttpcontroller.New(useCases)

	err := controller.UpdateUserAddress(req)
	require.ErrorIs(err, expectedErr)
}

func TestUpdateUserAddressUseCaseFail(t *testing.T) {
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

	body, err := userent.NewRegisteredAddress(addressID.String(), address)
	require.Nil(err)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}
	useCases.
		On("UpdateUserAddress", ctx, userID, body).
		Return(expectedErr)

	req := &mocks.Request{}
	req.On("ParseBody",
		mock.AnythingOfType("*userent.RegisteredAddress"),
	).Return(nil).Run(func(args mock.Arguments) {
		argAddress := args.Get(0).(*userent.RegisteredAddress)
		*argAddress = body
	})
	req.On("Context").Return(ctx)
	req.On("GetUserID").Return(userID, nil)

	controller := userhttpcontroller.New(useCases)

	err = controller.UpdateUserAddress(req)
	require.ErrorIs(err, expectedErr)
}
