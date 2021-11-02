package userhttpcontroller_test

import (
	"context"
	"encoding/json"
	"errors"
	userent "ifoodish-store/services/user/domain/entity"
	uservo "ifoodish-store/services/user/domain/valueobject"
	userhttpcontroller "ifoodish-store/services/user/infrastructure/controller/http"
	"ifoodish-store/services/user/mocks"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisterUserSuccess(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()

	expectedUserID := uservo.GenerateNewUserID()

	password, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	passwordConfirm, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	user, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	body := userhttpcontroller.RegisterUserBody{
		Password:        password,
		PasswordConfirm: passwordConfirm,
		User:            user,
	}

	useCases := &mocks.UserUseCases{}
	useCases.
		On("RegisterUser", ctx, user, password, passwordConfirm).
		Return(expectedUserID, nil)

	req := &mocks.Request{}
	req.On("ParseBody",
		mock.AnythingOfType("*userhttpcontroller.RegisterUserBody"),
	).Return(nil).Run(func(args mock.Arguments) {
		argBody := args.Get(0).(*userhttpcontroller.RegisterUserBody)
		*argBody = body
	})
	req.On("Context").Return(ctx)

	controller := userhttpcontroller.New(useCases)

	userID, err := controller.RegisterUser(req)
	require.Nil(err)
	require.True(userID.Equals(expectedUserID))
}
func TestRegisterUserParseBodyFail(t *testing.T) {
	require := require.New(t)

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}

	req := &mocks.Request{}
	req.On("GetUserID").Return(uservo.GenerateNewUserID(), nil)
	req.On("ParseBody",
		mock.AnythingOfType("*userhttpcontroller.RegisterUserBody"),
	).Return(expectedErr)

	controller := userhttpcontroller.New(useCases)

	_, err := controller.RegisterUser(req)
	require.ErrorIs(err, expectedErr)
}

func TestRegisterUserUseCaseFail(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()

	password, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	passwordConfirm, err := uservo.NewPasswordRaw("321321")
	require.Nil(err)

	user, err := userent.NewUser(
		"João da Silva",
		"lala@lala.com",
		"24543211234",
	)
	require.Nil(err)

	body := userhttpcontroller.RegisterUserBody{
		Password:        password,
		PasswordConfirm: passwordConfirm,
		User:            user,
	}

	userID := uservo.GenerateNewUserID()

	expectedErr := errors.New("test error")

	useCases := &mocks.UserUseCases{}
	useCases.
		On("RegisterUser", ctx, user, password, passwordConfirm).
		Return(userID, expectedErr)

	req := &mocks.Request{}
	req.On("ParseBody",
		mock.AnythingOfType("*userhttpcontroller.RegisterUserBody"),
	).Return(nil).Run(func(args mock.Arguments) {
		argBody := args.Get(0).(*userhttpcontroller.RegisterUserBody)
		*argBody = body
	})
	req.On("Context").Return(ctx)

	controller := userhttpcontroller.New(useCases)

	_, err = controller.RegisterUser(req)
	require.ErrorIs(err, expectedErr)
}

func TestRegisterUserUnmarshalDomainFail(t *testing.T) {
	require := require.New(t)

	hostname := "@lala.com"

	for index, tc := range []struct {
		name            string
		email           string
		phone           string
		password        string
		passwordConfirm string
		err             error
	}{{
		name:            strings.Repeat("a", uservo.MaxUserNameLength),
		email:           strings.Repeat("a", uservo.MaxEmailLength-len(hostname)) + hostname,
		phone:           strings.Repeat("1", uservo.MaxPhoneLength),
		password:        strings.Repeat("1", uservo.MaxRawPasswordLength),
		passwordConfirm: strings.Repeat("1", uservo.MaxRawPasswordLength),
		err:             nil,
	}, {
		name:            strings.Repeat("a", uservo.MaxUserNameLength+1),
		email:           strings.Repeat("a", uservo.MaxEmailLength-len(hostname)) + hostname,
		phone:           strings.Repeat("1", uservo.MaxPhoneLength),
		password:        strings.Repeat("1", uservo.MaxRawPasswordLength),
		passwordConfirm: strings.Repeat("1", uservo.MaxRawPasswordLength),
		err:             uservo.ErrUserNameMaxLength,
	}, {
		name:            strings.Repeat("a", uservo.MaxUserNameLength),
		email:           strings.Repeat("a", uservo.MaxEmailLength-len(hostname)) + hostname,
		phone:           strings.Repeat("1", uservo.MaxPhoneLength),
		password:        strings.Repeat("1", uservo.MaxRawPasswordLength+1),
		passwordConfirm: strings.Repeat("1", uservo.MaxRawPasswordLength),
		err:             uservo.ErrRawPasswordMaxLength,
	}, {
		name:            strings.Repeat("a", uservo.MaxUserNameLength),
		email:           strings.Repeat("a", uservo.MaxEmailLength-len(hostname)) + hostname,
		phone:           strings.Repeat("1", uservo.MaxPhoneLength),
		password:        strings.Repeat("1", uservo.MaxRawPasswordLength),
		passwordConfirm: strings.Repeat("1", uservo.MaxRawPasswordLength+1),
		err:             uservo.ErrRawPasswordMaxLength,
	}} {
		b, err := json.Marshal(userhttpcontroller.RegisterUserBody{
			Password:        uservo.PasswordRaw(tc.password),
			PasswordConfirm: uservo.PasswordRaw(tc.passwordConfirm),
			User: userent.User{
				Name:  uservo.UserName(tc.name),
				Email: uservo.Email(tc.email),
				Phone: uservo.Phone(tc.phone),
			},
		})
		require.Nil(err)

		var body userhttpcontroller.RegisterUserBody
		err = body.UnmarshalJSON(b)
		require.ErrorIs(err, tc.err, "index %d", index)
	}
}

func TestRegisterUserUnmarshalFail(t *testing.T) {
	require := require.New(t)

	for index, tc := range []struct {
		data string
		err  error
	}{{
		data: `{
			"password":"321321",
			"password_confirm":"321321",
			"user":{
				"name":"Joao da silva",
				"email":"lala@lala.com"
				"phone": "24543211234",
			}
		}`,
		err: &json.SyntaxError{},
	}, {
		data: `{
			"password":"321321",
			"password_confirm":"321321",
			"user":{
				"name":"Joao da silva",
				"email":"lala@lala.com"
				"phone": 123123
			}
		}`,
		err: &json.UnmarshalTypeError{},
	}} {
		var body userhttpcontroller.RegisterUserBody
		err := json.Unmarshal([]byte(tc.data), &body)
		require.ErrorAs(err, &tc.err, "index %d", index)
	}
}
