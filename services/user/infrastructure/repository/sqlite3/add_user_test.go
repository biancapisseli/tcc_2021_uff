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

func TestAddUserSuccess(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}

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

	user, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	repo := userreposqlite3.New()

	userID, err := repo.AddUser(ctx, user, uservo.NewPasswordEncoded("password"))
	require.Nil(err)
	_, err = uservo.NewUserID(userID.String())
	require.Nil(err)
}

func TestAddUserForeignKeyFail(t *testing.T) {
	require := require.New(t)

	sqlResult := &mocks.SQLResult{}

	expectedErr := sqlite3.ErrConstraintUnique

	tx := &mocks.TransactionFinisher{}
	tx.
		On("NamedExec", mock.AnythingOfType("string"), mock.Anything).
		Return(sqlResult, expectedErr)

	db := &mocks.Transactioner{}
	db.On("BeginTransaction").Return(tx, nil)

	ctx, err := sqlxtx.BeginTransaction(db, context.Background())
	require.Nil(err)

	user, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	repo := userreposqlite3.New()

	userID, err := repo.AddUser(ctx, user, uservo.NewPasswordEncoded("password"))
	require.ErrorIs(err, expectedErr)
	_, err = uservo.NewUserID(userID.String())
	require.Nil(err)
}

func TestAddUserUnexpectedErrorFail(t *testing.T) {
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

	user, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	repo := userreposqlite3.New()

	userID, err := repo.AddUser(ctx, user, uservo.NewPasswordEncoded("password"))
	require.ErrorIs(err, expectedErr)
	_, err = uservo.NewUserID(userID.String())
	require.Nil(err)
}

func TestAddUserNoTransactionFail(t *testing.T) {
	require := require.New(t)

	expectedErr := sqlxtx.ErrTransactionNotFound

	user, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	repo := userreposqlite3.New()

	_, err = repo.AddUser(context.Background(), user, uservo.NewPasswordEncoded("password"))
	require.ErrorIs(err, expectedErr)
}
