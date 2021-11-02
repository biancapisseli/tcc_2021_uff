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

func TestGetUserInfoSuccess(t *testing.T) {
	require := require.New(t)

	userID := uservo.GenerateNewUserID()

	expectedUser, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	expectedRegAddress, err := userent.NewRegisteredUser(
		userID.String(),
		expectedUser,
	)
	require.Nil(err)

	tx := &mocks.TransactionFinisher{}
	tx.
		On("Get",
			mock.AnythingOfType("*userent.RegisteredUser"),
			mock.AnythingOfType("string"),
			userID,
		).
		Return(nil).
		Run(func(args mock.Arguments) {
			argAddress := args.Get(0).(*userent.RegisteredUser)
			*argAddress = expectedRegAddress
		})

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	user, err := repo.GetUserInfo(ctx, userID)
	require.Nil(err)
	require.True(user.ID.Equals(userID))
}

func TestGetUserInfoNoRowsFail(t *testing.T) {
	require := require.New(t)

	userID := uservo.GenerateNewUserID()

	expectedErr := sql.ErrNoRows

	tx := &mocks.TransactionFinisher{}
	tx.
		On("Get",
			mock.AnythingOfType("*userent.RegisteredUser"),
			mock.AnythingOfType("string"),
			userID,
		).
		Return(expectedErr)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	_, err = repo.GetUserInfo(ctx, userID)
	require.ErrorIs(err, expectedErr)
}

func TestGetUserInfoUnexpectedErrorFail(t *testing.T) {
	require := require.New(t)

	userID := uservo.GenerateNewUserID()

	expectedErr := errors.New("test-error")

	tx := &mocks.TransactionFinisher{}
	tx.
		On("Get",
			mock.AnythingOfType("*userent.RegisteredUser"),
			mock.AnythingOfType("string"),
			userID,
		).
		Return(expectedErr)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	_, err = repo.GetUserInfo(ctx, userID)
	require.ErrorIs(err, expectedErr)
}

func TestGetUserInfoNoTransactionFail(t *testing.T) {
	require := require.New(t)

	expectedErr := sqlxtx.ErrTransactionNotFound

	userID := uservo.GenerateNewUserID()

	repo := userreposqlite3.New()

	_, err := repo.GetUserInfo(context.Background(), userID)
	require.ErrorIs(err, expectedErr)
}
