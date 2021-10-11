package userdom

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	validUser = User{
		Name:  UserName((strings.Repeat("a", maxUserNameLength-1))),
		Email: Email((strings.Repeat("a", maxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME)),
	}
	invalidUser1 = User{
		Name:  UserName((strings.Repeat("a", maxUserNameLength+1))),
		Email: Email((strings.Repeat("a", maxEmailLength))),
	}
	invalidUser2 = User{
		Name:  UserName((strings.Repeat("a", maxUserNameLength))),
		Email: Email((strings.Repeat("a", maxEmailLength+1))),
	}
	validRegisteredUser = RegisteredUser{
		User: validUser,
		ID:   50,
	}
	invalidRegisteredUser1 = RegisteredUser{
		User: invalidUser1,
		ID:   -50,
	}
	invalidRegisteredUser2 = RegisteredUser{
		User: validUser,
		ID:   -50,
	}
)

func TestUserValid(t *testing.T) {
	require := require.New(t)

	user, myError := NewUser(validUser)
	require.Nil(myError)
	require.NotEmpty(user)

}

func TestUserInalid(t *testing.T) {
	require := require.New(t)

	user, myError := NewUser(invalidUser1)
	require.NotNil(myError)
	require.Nil(user)

	user, myError = NewUser(invalidUser2)
	require.NotNil(myError)
	require.Nil(user)

}

func TestRegisteredUserValid(t *testing.T) {
	require := require.New(t)

	user, myError := NewRegisteredUser(validRegisteredUser)
	require.Nil(myError)
	require.NotEmpty(user)

}

func TestRegisteredUserInvalid(t *testing.T) {
	require := require.New(t)

	user, myError := NewRegisteredUser(invalidRegisteredUser1)
	require.NotNil(myError)
	require.Nil(user)

	user, myError = NewRegisteredUser(invalidRegisteredUser2)
	require.NotNil(myError)
	require.Nil(user)
}
func TestJSONUnmarshallingSuccess(t *testing.T) {
	require := require.New(t)

	var user *User
	err := json.Unmarshal([]byte(`
		{
			"name":"Matheus Zabin",
			"email":"lalal@lala.com"
		}
	`), &user)
	require.Nil(err)
	require.True(user.Name.Equals("Matheus Zabin"))
	require.True(user.Email.Equals("lalal@lala.com"))

}

func TestJSONUnmarshallingFail(t *testing.T) {
	require := require.New(t)
	var user *User

	// forçando teste do unmarshal
	err := user.UnmarshalJSON([]byte(`
		{
			"name":"Matheus Zabin",
			"email":"lalal@lala.com
		}
	`))
	require.NotNil(err)

	user = nil
	err = json.Unmarshal([]byte(`
		{
			"name":"Aa",
			"email":"lalal@lala.com"
		}
	`), &user)
	require.ErrorIs(err, ErrUserNameMinLength)

	user = nil
	err = json.Unmarshal([]byte(`
		{
			"name":"Matheus Zabin",
			"email":"emailInvalido"
		}
	`), &user)
	require.ErrorIs(err, ErrEmailInvalidFormat)

}
