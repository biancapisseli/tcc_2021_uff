package userent_test

import (
	"encoding/json"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"

	"github.com/carlmjohnson/resperr"

	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

const (
	EMAIL_HOSTMAME = "@example.com"
)

type userTestCase struct {
	name  string
	email string
	phone string

	expectedErr error
}

var (
	validUserTestCase = userTestCase{
		name:  "João da Silva",
		email: strings.Repeat("a", uservo.MaxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME,
		phone: "24999999999",
	}
)

func userTestCaseCompare(
	require *require.Assertions,
	user userent.User,
	tc userTestCase,
) {
	require.Equal(tc.name, user.Name.String())
	require.Equal(tc.email, user.Email.String())
	require.Equal(tc.phone, user.Phone.String())
}

func TestUserValid(t *testing.T) {
	require := require.New(t)

	user, err := userent.NewUser(
		validUserTestCase.name,
		validUserTestCase.email,
		validUserTestCase.phone,
	)
	require.Nil(err)
	userTestCaseCompare(require, user, validUserTestCase)
}

func TestUserInvalid(t *testing.T) {
	require := require.New(t)

	users := []userTestCase{{
		name:  "João 123",
		email: validUserTestCase.email,
		phone: validUserTestCase.phone,

		expectedErr: uservo.ErrUserNameInvalidCharacter,
	}, {
		name:  strings.Repeat("a", uservo.MaxUserNameLength+1),
		email: validUserTestCase.email,
		phone: validUserTestCase.phone,

		expectedErr: uservo.ErrUserNameMaxLength,
	}, {
		name:  strings.Repeat("a", uservo.MinUserNameLength-1),
		email: validUserTestCase.email,
		phone: validUserTestCase.phone,

		expectedErr: uservo.ErrUserNameMinLength,
	}, {
		name:  validUserTestCase.name,
		email: "João123",
		phone: validUserTestCase.phone,

		expectedErr: uservo.ErrEmailInvalidFormat,
	}, {
		name:  validUserTestCase.name,
		email: strings.Repeat("a", uservo.MaxEmailLength) + EMAIL_HOSTMAME,
		phone: validUserTestCase.phone,

		expectedErr: uservo.ErrEmailMaxLength,
	}, {
		name:  validUserTestCase.name,
		email: validUserTestCase.email,
		phone: strings.Repeat("9", uservo.MaxPhoneLength+1),

		expectedErr: uservo.ErrPhoneMaxLength,
	}, {
		name:  validUserTestCase.name,
		email: validUserTestCase.email,
		phone: strings.Repeat("a", uservo.MinPhoneLength-1),

		expectedErr: uservo.ErrPhoneMinLength,
	}}

	for i, it := range users {
		_, err := userent.NewUser(
			it.name,
			it.email,
			it.phone,
		)
		require.ErrorIs(err, it.expectedErr, "index %d", i)
	}
}

func TestRegisteredUserValid(t *testing.T) {
	require := require.New(t)

	user, err := userent.NewUser(
		validUserTestCase.name,
		validUserTestCase.email,
		validUserTestCase.phone,
	)
	require.Nil(err)

	userUUID := uuid.New().String()

	regUser, err := userent.NewRegisteredUser(userUUID, user)
	require.Nil(err)
	userTestCaseCompare(require, regUser.User, validUserTestCase)
	require.Equal(userUUID, regUser.ID.String())
}

func TestRegisteredUserInvalid(t *testing.T) {
	require := require.New(t)

	user, err := userent.NewUser(
		validUserTestCase.name,
		validUserTestCase.email,
		validUserTestCase.phone,
	)
	require.Nil(err)

	_, err = userent.NewRegisteredUser(uuid.Nil.String(), user)
	require.Equal(http.StatusBadRequest, resperr.StatusCode(err))
}

func TestJSONUnmarshallingUserSuccess(t *testing.T) {
	require := require.New(t)

	var user *userent.User
	err := json.Unmarshal([]byte(`
		{
			"name":"Matheus Zabin",
			"email":"lalal@lala.com",
			"phone":"24999999999"
		}
	`), &user)
	require.Nil(err)
	require.True(user.Name.Equals("Matheus Zabin"))
	require.True(user.Email.Equals("lalal@lala.com"))
	require.True(user.Phone.Equals("24999999999"))
}

func TestJSONUnmarshallingUserFail(t *testing.T) {
	require := require.New(t)
	var user *userent.User

	// forçando teste do unmarshal
	err := user.UnmarshalJSON([]byte(`
		{
			"name":"Matheus Zabin",
			"email":"lalal@lala.com
			"phone":"249999224073"
		}
	`))
	require.NotNil(err)

	user = nil
	err = json.Unmarshal([]byte(`
		{
			"name":"Aa",
			"email":"lalal@lala.com",
			"phone":"249999224073"
		}
	`), &user)
	require.ErrorIs(err, uservo.ErrUserNameMinLength)

	user = nil
	err = json.Unmarshal([]byte(`
		{
			"name":"Matheus Zabin",
			"email":"emailInvalido",
			"phone":"249999224073"
		}
	`), &user)
	require.ErrorIs(err, uservo.ErrEmailInvalidFormat)

	user = nil
	err = json.Unmarshal([]byte(`
		{
			"name":"Matheus Zabin",
			"email":"lalal@lala.com",
			"phone":"249993"
		}
	`), &user)
	require.ErrorIs(err, uservo.ErrPhoneMinLength)

}

///////////////////////////////////////////////
func TestJSONUnmarshallingRegisteredUserSuccess(t *testing.T) {
	require := require.New(t)

	var registeredUser userent.RegisteredUser
	err := registeredUser.UnmarshalJSON([]byte(`
	{
		"id":"123e4567-e89b-12d3-a456-426614174000",
		"name":"Matheus Zabin",
		"email":"lalal@lala.com",
		"phone":"24999999999"
	}
	`))
	require.Nil(err)
	require.True(registeredUser.ID.Equals(uservo.UserID(uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"))))
	require.True(registeredUser.Name.Equals("Matheus Zabin"))
	require.True(registeredUser.Email.Equals("lalal@lala.com"))
	require.True(registeredUser.Phone.Equals("24999999999"))
}

func TestJSONUnmarshallingRegisteredUserFail(t *testing.T) {
	require := require.New(t)
	var registeredUser userent.RegisteredUser

	// forçando teste do unmarshal
	err := registeredUser.UnmarshalJSON([]byte(`
		{
			"id":"123e4567-e89b-12d3-a456-426614174000",
			"name":"Matheus Zabin",
			"email":"lalal@lala.com
			"phone":"249999224073"
		}
	`))
	require.Equal(http.StatusInternalServerError, resperr.StatusCode(err))

	err = json.Unmarshal([]byte(`
		{
			"id":"123e4567-e89b-12d3-a456-426614174000",
			"name":"Aa",
			"email":"lalal@lala.com",
			"phone":"249999224073"
		}
	`), &registeredUser)
	require.ErrorIs(err, uservo.ErrUserNameMinLength)

	err = json.Unmarshal([]byte(`
		{
			"id":"123e4567-e89b-12d3-a456-426614174000",
			"name":"Matheus Zabin",
			"email":"emailInvalido",
			"phone":"24999999999"
		}
	`), &registeredUser)
	require.ErrorIs(err, uservo.ErrEmailInvalidFormat)

	err = json.Unmarshal([]byte(`
		{
			"id":"123e4567-e89b-12d3-a456-426614174000",
			"name":"Matheus Zabin",
			"email":"lalal@lala.com",
			"phone":"249993"
		}
	`), &registeredUser)
	require.ErrorIs(err, uservo.ErrPhoneMinLength)

	err = json.Unmarshal([]byte(`
		{
			"id":"macarrao",
			"name":"Matheus Zabin",
			"email":"lalal@lala.com",
			"phone":"24999999999"
		}
	`), &registeredUser)
	require.Equal(http.StatusBadRequest, resperr.StatusCode(err))

	err = json.Unmarshal([]byte(`
		{
			"id": 10.24,
			"name":"Matheus Zabin",
			"email":"lalal@lala.com",
			"phone":"24999999999"
		}
	`), &registeredUser)
	require.Equal(http.StatusBadRequest, resperr.StatusCode(err))

	err = json.Unmarshal([]byte(`
		{
			"id": "00000000-0000-0000-0000-000000000000",
			"name":"Matheus Zabin",
			"email":"lalal@lala.com",
			"phone":"24999999999"
		}
	`), &registeredUser)
	require.Equal(http.StatusBadRequest, resperr.StatusCode(err))

}
