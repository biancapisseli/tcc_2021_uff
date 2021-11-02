package userreposqlite3_test

import (
	"context"
	"errors"
	"ifoodish-store/mocks"
	"ifoodish-store/pkg/sqlxtx"
	"net/http"
	"testing"

	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
	userreposqlite3 "ifoodish-store/services/user/infrastructure/repository/sqlite3"

	"github.com/carlmjohnson/resperr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSaveUserSuccess(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}
	sqlResult.On("RowsAffected").Return(int64(1), nil)

	userID := uservo.GenerateNewUserID()

	user, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	tx := &mocks.TransactionFinisher{}
	tx.
		On("NamedExec",
			mock.AnythingOfType("string"),
			mock.Anything,
		).
		Return(sqlResult, nil)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	err = repo.SaveUser(ctx, userID, user)
	require.Nil(err)
}

func TestSaveUserUnexpectedErrorFail(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}

	userID := uservo.GenerateNewUserID()

	user, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	expectedErr := errors.New("test-error")

	tx := &mocks.TransactionFinisher{}
	tx.
		On("NamedExec",
			mock.AnythingOfType("string"),
			mock.Anything,
		).
		Return(sqlResult, expectedErr)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	err = repo.SaveUser(ctx, userID, user)
	require.ErrorIs(err, expectedErr)
}

func TestSaveUserRowsAffectedErrorFail(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}
	sqlResult.On("RowsAffected").Return(int64(0), errors.New("test-error"))

	userID := uservo.GenerateNewUserID()

	user, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	tx := &mocks.TransactionFinisher{}
	tx.
		On("NamedExec",
			mock.AnythingOfType("string"),
			mock.Anything,
		).
		Return(sqlResult, nil)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	err = repo.SaveUser(ctx, userID, user)
	require.Equal(http.StatusInternalServerError, resperr.StatusCode(err))
}

func TestSaveUserZeroRowsAffectedFail(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}
	sqlResult.On("RowsAffected").Return(int64(0), nil)

	userID := uservo.GenerateNewUserID()

	user, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	tx := &mocks.TransactionFinisher{}
	tx.
		On("NamedExec",
			mock.AnythingOfType("string"),
			mock.Anything,
		).
		Return(sqlResult, nil)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	repo := userreposqlite3.New()

	err = repo.SaveUser(ctx, userID, user)
	require.Equal(http.StatusNotFound, resperr.StatusCode(err))
}

func TestSaveUserNoTransactionFail(t *testing.T) {
	require := require.New(t)

	expectedErr := sqlxtx.ErrTransactionNotFound

	userID := uservo.GenerateNewUserID()

	user, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	repo := userreposqlite3.New()

	err = repo.SaveUser(context.Background(), userID, user)
	require.ErrorIs(err, expectedErr)
}
