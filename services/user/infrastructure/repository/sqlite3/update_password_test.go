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

func TestUpdatePasswordSuccess(t *testing.T) {
	require := require.New(t)

	password := uservo.NewPasswordEncoded("password")

	sqlResult := &mocks.SQLResult{}
	sqlResult.On("RowsAffected").Return(int64(1), nil)

	userID := uservo.GenerateNewUserID()

	tx := &mocks.TransactionFinisher{}
	tx.
		On("Exec",
			mock.AnythingOfType("string"),
			password,
			userID,
		).
		Return(sqlResult, nil)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	err = repo.UpdatePassword(ctx, userID, password)
	require.Nil(err)
}

func TestUpdatePasswordUnexpectedErrorFail(t *testing.T) {
	require := require.New(t)

	password := uservo.NewPasswordEncoded("password")

	sqlResult := &mocks.SQLResult{}

	userID := uservo.GenerateNewUserID()

	expectedErr := errors.New("test-error")

	tx := &mocks.TransactionFinisher{}
	tx.
		On("Exec",
			mock.AnythingOfType("string"),
			password,
			userID,
		).
		Return(sqlResult, expectedErr)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	err = repo.UpdatePassword(ctx, userID, password)
	require.ErrorIs(err, expectedErr)
}

func TestUpdatePasswordRowsAffectedErrorFail(t *testing.T) {
	require := require.New(t)

	password := uservo.NewPasswordEncoded("password")

	sqlResult := &mocks.SQLResult{}
	sqlResult.On("RowsAffected").Return(int64(0), errors.New("test-error"))

	userID := uservo.GenerateNewUserID()

	tx := &mocks.TransactionFinisher{}
	tx.
		On("Exec",
			mock.AnythingOfType("string"),
			password,
			userID,
		).
		Return(sqlResult, nil)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	err = repo.UpdatePassword(ctx, userID, password)
	require.Equal(http.StatusInternalServerError, resperr.StatusCode(err))
}

func TestUpdatePasswordZeroRowsAffectedFail(t *testing.T) {
	require := require.New(t)

	password := uservo.NewPasswordEncoded("password")

	sqlResult := &mocks.SQLResult{}
	sqlResult.On("RowsAffected").Return(int64(0), nil)

	userID := uservo.GenerateNewUserID()

	tx := &mocks.TransactionFinisher{}
	tx.
		On("Exec",
			mock.AnythingOfType("string"),
			password,
			userID,
		).
		Return(sqlResult, nil)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	err = repo.UpdatePassword(ctx, userID, password)
	require.Equal(http.StatusNotFound, resperr.StatusCode(err))
}

func TestUpdatePasswordNoTransactionFail(t *testing.T) {
	require := require.New(t)

	password := uservo.NewPasswordEncoded("password")

	expectedErr := sqlxtx.ErrTransactionNotFound

	userID := uservo.GenerateNewUserID()

	repo := userreposqlite3.New()

	err := repo.UpdatePassword(context.Background(), userID, password)
	require.ErrorIs(err, expectedErr)
}
