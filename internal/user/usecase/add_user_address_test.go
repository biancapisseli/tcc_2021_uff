package useruc_test

import (
	"context"
	"testing"

	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/internal/user/mocks"
	useruc "ifoodish-store/internal/user/usecase"

	"github.com/stretchr/testify/require"
)

func TestAddUserAddressSuccess(t *testing.T) {
	require := require.New(t)

	repo := &mocks.UserRepository{}
	encoder := &mocks.PasswordEncoder{}

	useCases := useruc.New(repo, encoder)

	ctx := context.Background()

	userID := uservo.GenerateNewUserID()

	useCases.AddUserAddress(ctx)
}
