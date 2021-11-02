package userreposqlite3_test

import (
	"context"
	"errors"
	"ifoodish-store/mocks"
	"ifoodish-store/pkg/sqlxtx"
	"net/http"
	"testing"

	uservo "ifoodish-store/services/user/domain/valueobject"
	userreposqlite3 "ifoodish-store/services/user/infrastructure/repository/sqlite3"

	"github.com/carlmjohnson/resperr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRemoveUserAddressSuccess(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}
	sqlResult.On("RowsAffected").Return(int64(1), nil)

	userID := uservo.GenerateNewUserID()

	addressID := uservo.GenerateNewAddressID()

	tx := &mocks.TransactionFinisher{}
	tx.
		On("Exec",
			mock.AnythingOfType("string"),
			addressID,
			userID,
		).
		Return(sqlResult, nil)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	err = repo.RemoveUserAddress(ctx, userID, addressID)
	require.Nil(err)
}

func TestRemoveUserAddressUnexpectedErrorFail(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}

	userID := uservo.GenerateNewUserID()

	addressID := uservo.GenerateNewAddressID()

	expectedErr := errors.New("test-error")

	tx := &mocks.TransactionFinisher{}
	tx.
		On("Exec",
			mock.AnythingOfType("string"),
			addressID,
			userID,
		).
		Return(sqlResult, expectedErr)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	err = repo.RemoveUserAddress(ctx, userID, addressID)
	require.ErrorIs(err, expectedErr)
}

func TestRemoveUserAddressRowsAffectedErrorFail(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}
	sqlResult.On("RowsAffected").Return(int64(0), errors.New("test-error"))

	userID := uservo.GenerateNewUserID()

	addressID := uservo.GenerateNewAddressID()

	tx := &mocks.TransactionFinisher{}
	tx.
		On("Exec",
			mock.AnythingOfType("string"),
			addressID,
			userID,
		).
		Return(sqlResult, nil)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	err = repo.RemoveUserAddress(ctx, userID, addressID)
	require.Equal(http.StatusInternalServerError, resperr.StatusCode(err))
}

func TestRemoveUserAddressZeroRowsAffectedFail(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}
	sqlResult.On("RowsAffected").Return(int64(0), nil)

	userID := uservo.GenerateNewUserID()

	addressID := uservo.GenerateNewAddressID()

	tx := &mocks.TransactionFinisher{}
	tx.
		On("Exec",
			mock.AnythingOfType("string"),
			addressID,
			userID,
		).
		Return(sqlResult, nil)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	err = repo.RemoveUserAddress(ctx, userID, addressID)
	require.Equal(http.StatusNotFound, resperr.StatusCode(err))
}

func TestRemoveUserAddressNoTransactionFail(t *testing.T) {
	require := require.New(t)

	expectedErr := sqlxtx.ErrTransactionNotFound

	userID := uservo.GenerateNewUserID()
	addressID := uservo.GenerateNewAddressID()

	repo := userreposqlite3.New()

	err := repo.RemoveUserAddress(context.Background(), userID, addressID)
	require.ErrorIs(err, expectedErr)
}
