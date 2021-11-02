package userreposqlite3_test

import (
	"context"
	"database/sql"
	"errors"
	"ifoodish-store/mocks"
	"ifoodish-store/pkg/sqlxtx"
	"testing"

	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
	userreposqlite3 "ifoodish-store/services/user/infrastructure/repository/sqlite3"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetUserAddressSuccess(t *testing.T) {
	require := require.New(t)

	userID := uservo.GenerateNewUserID()
	addressID := uservo.GenerateNewAddressID()

	expectedAddress, err := userent.NewAddress(
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

	expectedRegAddress, err := userent.NewRegisteredAddress(
		addressID.String(),
		expectedAddress,
	)
	require.Nil(err)

	tx := &mocks.TransactionFinisher{}
	tx.
		On("Get",
			mock.AnythingOfType("*userent.RegisteredAddress"),
			mock.AnythingOfType("string"),
			userID,
			addressID,
		).
		Return(nil).
		Run(func(args mock.Arguments) {
			argAddress := args.Get(0).(*userent.RegisteredAddress)
			*argAddress = expectedRegAddress
		})

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	address, err := repo.GetUserAddress(ctx, userID, addressID)
	require.Nil(err)
	require.True(address.ID.Equals(addressID))
}

func TestGetUserAddressNoRowsFail(t *testing.T) {
	require := require.New(t)

	userID := uservo.GenerateNewUserID()
	addressID := uservo.GenerateNewAddressID()

	expectedErr := sql.ErrNoRows

	tx := &mocks.TransactionFinisher{}
	tx.
		On("Get",
			mock.AnythingOfType("*userent.RegisteredAddress"),
			mock.AnythingOfType("string"),
			userID,
			addressID,
		).
		Return(expectedErr)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	_, err = repo.GetUserAddress(ctx, userID, addressID)
	require.ErrorIs(err, expectedErr)
}

func TestGetUserAddressUnexpectedErrorFail(t *testing.T) {
	require := require.New(t)

	userID := uservo.GenerateNewUserID()
	addressID := uservo.GenerateNewAddressID()

	expectedErr := errors.New("test-error")

	tx := &mocks.TransactionFinisher{}
	tx.
		On("Get",
			mock.AnythingOfType("*userent.RegisteredAddress"),
			mock.AnythingOfType("string"),
			userID,
			addressID,
		).
		Return(expectedErr)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	_, err = repo.GetUserAddress(ctx, userID, addressID)
	require.ErrorIs(err, expectedErr)
}

func TestGetUserAddressNoTransactionFail(t *testing.T) {
	require := require.New(t)

	expectedErr := sqlxtx.ErrTransactionNotFound

	userID := uservo.GenerateNewUserID()
	addressID := uservo.GenerateNewAddressID()

	repo := userreposqlite3.New()

	_, err := repo.GetUserAddress(context.Background(), userID, addressID)
	require.ErrorIs(err, expectedErr)
}
