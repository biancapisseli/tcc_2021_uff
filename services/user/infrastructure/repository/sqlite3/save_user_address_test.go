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

func TestSaveUserAddressSuccess(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}
	sqlResult.On("RowsAffected").Return(int64(1), nil)

	userID := uservo.GenerateNewUserID()

	addressID := uservo.GenerateNewAddressID()

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

	regAddress, err := userent.NewRegisteredAddress(
		addressID.String(),
		address,
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

	err = repo.SaveUserAddress(ctx, userID, regAddress)
	require.Nil(err)
}

func TestSaveUserAddressUnexpectedErrorFail(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}

	userID := uservo.GenerateNewUserID()

	addressID := uservo.GenerateNewAddressID()

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

	regAddress, err := userent.NewRegisteredAddress(
		addressID.String(),
		address,
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

	err = repo.SaveUserAddress(ctx, userID, regAddress)
	require.ErrorIs(err, expectedErr)
}

func TestSaveUserAddressRowsAffectedErrorFail(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}
	sqlResult.On("RowsAffected").Return(int64(0), errors.New("test-error"))

	userID := uservo.GenerateNewUserID()

	addressID := uservo.GenerateNewAddressID()

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

	regAddress, err := userent.NewRegisteredAddress(
		addressID.String(),
		address,
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

	err = repo.SaveUserAddress(ctx, userID, regAddress)
	require.Equal(http.StatusInternalServerError, resperr.StatusCode(err))
}

func TestSaveUserAddressZeroRowsAffectedFail(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}
	sqlResult.On("RowsAffected").Return(int64(0), nil)

	userID := uservo.GenerateNewUserID()

	addressID := uservo.GenerateNewAddressID()

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

	regAddress, err := userent.NewRegisteredAddress(
		addressID.String(),
		address,
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

	err = repo.SaveUserAddress(ctx, userID, regAddress)
	require.Equal(http.StatusNotFound, resperr.StatusCode(err))
}

func TestSaveUserAddressNoTransactionFail(t *testing.T) {
	require := require.New(t)

	expectedErr := sqlxtx.ErrTransactionNotFound

	userID := uservo.GenerateNewUserID()

	addressID := uservo.GenerateNewAddressID()

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

	regAddress, err := userent.NewRegisteredAddress(
		addressID.String(),
		address,
	)
	require.Nil(err)

	repo := userreposqlite3.New()

	err = repo.SaveUserAddress(context.Background(), userID, regAddress)
	require.ErrorIs(err, expectedErr)
}
