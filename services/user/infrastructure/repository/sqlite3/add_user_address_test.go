package userreposqlite3_test

import (
	"context"
	"errors"
	"ifoodish-store/mocks"
	"ifoodish-store/pkg/sqlxtx"
	"testing"

	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
	userreposqlite3 "ifoodish-store/services/user/infrastructure/repository/sqlite3"

	"github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAddUserAddressSuccess(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}

	tx := &mocks.TransactionFinisher{}
	tx.
		On("NamedExec", mock.AnythingOfType("string"), mock.Anything).
		Return(sqlResult, nil)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

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

	repo := userreposqlite3.New()

	addressID, err := repo.AddUserAddress(ctx, userID, address)
	require.Nil(err)
	_, err = uservo.NewAddressID(addressID.String())
	require.Nil(err)
}

func TestAddUserAddressForeignKeyFail(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}

	expectedErr := sqlite3.ErrConstraintForeignKey

	tx := &mocks.TransactionFinisher{}
	tx.
		On("NamedExec", mock.AnythingOfType("string"), mock.Anything).
		Return(sqlResult, expectedErr)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

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

	repo := userreposqlite3.New()

	addressID, err := repo.AddUserAddress(ctx, userID, address)
	require.ErrorIs(err, expectedErr)
	_, err = uservo.NewAddressID(addressID.String())
	require.Nil(err)
}

func TestAddUserAddressUnexpectedErrorFail(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}

	expectedErr := errors.New("test error")

	tx := &mocks.TransactionFinisher{}
	tx.
		On("NamedExec", mock.AnythingOfType("string"), mock.Anything).
		Return(sqlResult, expectedErr)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

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

	repo := userreposqlite3.New()

	addressID, err := repo.AddUserAddress(ctx, userID, address)
	require.ErrorIs(err, expectedErr)
	_, err = uservo.NewAddressID(addressID.String())
	require.Nil(err)
}

func TestAddUserAddressNoTransactionFail(t *testing.T) {
	require := require.New(t)

	expectedErr := sqlxtx.ErrTransactionNotFound

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

	repo := userreposqlite3.New()

	_, err = repo.AddUserAddress(context.Background(), userID, address)
	require.ErrorIs(err, expectedErr)
}
