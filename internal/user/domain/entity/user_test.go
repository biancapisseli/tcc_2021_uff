package userent_test

import (
	"encoding/json"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/resperr"

	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

const (
	EMAIL_HOSTMAME = "@example.com"
)

var (
	validUser = userent.User{
		Name:  "João da Silva",
		Email: uservo.Email(strings.Repeat("a", uservo.MaxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME),
		Phone: "24999999999",
	}
)

func TestUserValid(t *testing.T) {
	require := require.New(t)

	user, err := userent.NewUser(validUser)
	require.Nil(err)
	require.NotEmpty(user)
}

func TestUserInvalid(t *testing.T) {
	require := require.New(t)

	type testIterator struct {
		user userent.User
		err  error
	}

	users := []testIterator{}

	//Name
	example := validUser
	example.Name = uservo.UserName("João 123")
	users = append(users, testIterator{
		user: example,
		err:  uservo.ErrUserNameInvalidCharacter,
	})

	example = validUser
	example.Name = uservo.UserName(strings.Repeat("a", uservo.MaxUserNameLength+1))
	users = append(users, testIterator{
		user: example,
		err:  uservo.ErrUserNameMaxLength,
	})

	example = validUser
	example.Name = uservo.UserName(strings.Repeat("a", uservo.MinUserNameLength-1))
	users = append(users, testIterator{
		user: example,
		err:  uservo.ErrUserNameMinLength,
	})

	//Email
	example = validUser
	example.Email = uservo.Email("João123")
	users = append(users, testIterator{
		user: example,
		err:  uservo.ErrEmailInvalidFormat,
	})

	example = validUser
	example.Email = uservo.Email(strings.Repeat("a", uservo.MaxEmailLength) + EMAIL_HOSTMAME)
	users = append(users, testIterator{
		user: example,
		err:  uservo.ErrEmailMaxLength,
	})

	//Phone
	example = validUser
	example.Phone = uservo.Phone(strings.Repeat("9", uservo.MaxPhoneLength+1))
	users = append(users, testIterator{
		user: example,
		err:  uservo.ErrPhoneMaxLength,
	})

	example = validUser
	example.Phone = uservo.Phone(strings.Repeat("a", uservo.MinPhoneLength-1))
	users = append(users, testIterator{
		user: example,
		err:  uservo.ErrPhoneMinLength,
	})

	for _, it := range users {
		newUser, err := userent.NewUser(it.user)
		require.ErrorIs(err, it.err)
		require.Nil(newUser)
	}

}

func TestRegisteredUserValid(t *testing.T) {
	require := require.New(t)

	ex := userent.RegisteredUser{}
	ex.User = validUser
	ex.ID = uservo.UserID(uuid.New())

	user, err := userent.NewRegisteredUser(ex)
	require.Nil(err)
	require.NotEmpty(user)

}

func TestRegisteredUserInvalid(t *testing.T) {
	require := require.New(t)

	ex := userent.RegisteredUser{}
	ex.User = validUser

	ex.ID = uservo.UserID([16]byte{0, 0})
	_, err := userent.NewRegisteredUser(ex)
	require.Equal(http.StatusBadRequest, resperr.StatusCode(err))

	ex.Name = uservo.UserName(strings.Repeat("a", uservo.MinUserNameLength-1))
	_, err = userent.NewRegisteredUser(ex)
	require.ErrorIs(err, uservo.ErrUserNameMinLength)

}

func TestJSONUnmarshallingSuccess(t *testing.T) {
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

func TestJSONUnmarshallingFail(t *testing.T) {
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
