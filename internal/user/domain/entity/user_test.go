package userent_test

// const (
// 	EMAIL_HOSTMAME = "@example.com"
// )

// var (
// 	validUser = userent.User{
// 		Name:  strings.Repeat("a", uservo.MaxUserNameLength-1)),
// 		Email: strings.Repeat("a", uservo.MaxEmailLength-len(EMAIL_HOSTMAME)) + EMAIL_HOSTMAME),
// 	}
// 	invalidUser1 = User{
// 		Name:  strings.Repeat("a", uservo.MaxUserNameLength+1),
// 		Email: strings.Repeat("a", uservo.MaxEmailLength),
// 	}
// 	invalidUser2 = User{
// 		Name:  strings.Repeat("a", uservo.MaxUserNameLength)),
// 		Email: strings.Repeat("a", uservo.MaxEmailLength+1)),
// 	}
// 	validRegisteredUser = RegisteredUser{
// 		User: validUser,
// 		ID:   50,
// 	}
// 	invalidRegisteredUser1 = RegisteredUser{
// 		User: invalidUser1,
// 		ID:   -50,
// 	}
// 	invalidRegisteredUser2 = RegisteredUser{
// 		User: validUser,
// 		ID:   -50,
// 	}
// )

// func TestUserValid(t *testing.T) {
// 	require := require.New(t)

// 	user, myError := NewUser(validUser)
// 	require.Nil(myError)
// 	require.NotEmpty(user)

// }

// func TestUserInvalid(t *testing.T) {
// 	require := require.New(t)

// 	user, myError := NewUser(invalidUser1)
// 	require.NotNil(myError)
// 	require.Nil(user)

// 	user, myError = NewUser(invalidUser2)
// 	require.NotNil(myError)
// 	require.Nil(user)

// }

// func TestRegisteredUserValid(t *testing.T) {
// 	require := require.New(t)

// 	user, myError := NewRegisteredUser(validRegisteredUser)
// 	require.Nil(myError)
// 	require.NotEmpty(user)

// }

// func TestRegisteredUserInvalid(t *testing.T) {
// 	require := require.New(t)

// 	user, myError := NewRegisteredUser(invalidRegisteredUser1)
// 	require.NotNil(myError)
// 	require.Nil(user)

// 	user, myError = NewRegisteredUser(invalidRegisteredUser2)
// 	require.NotNil(myError)
// 	require.Nil(user)
// }
// func TestJSONUnmarshallingSuccess(t *testing.T) {
// 	require := require.New(t)

// 	var user *User
// 	err := json.Unmarshal([]byte(`
// 		{
// 			"name":"Matheus Zabin",
// 			"email":"lalal@lala.com"
// 		}
// 	`), &user)
// 	require.Nil(err)
// 	require.True(user.Name.Equals("Matheus Zabin"))
// 	require.True(user.Email.Equals("lalal@lala.com"))

// }

// func TestJSONUnmarshallingFail(t *testing.T) {
// 	require := require.New(t)
// 	var user *User

// 	// forçando teste do unmarshal
// 	err := user.UnmarshalJSON([]byte(`
// 		{
// 			"name":"Matheus Zabin",
// 			"email":"lalal@lala.com
// 		}
// 	`))
// 	require.NotNil(err)

// 	user = nil
// 	err = json.Unmarshal([]byte(`
// 		{
// 			"name":"Aa",
// 			"email":"lalal@lala.com"
// 		}
// 	`), &user)
// 	require.ErrorIs(err, uservo.ErrUserNameMinLength)

// 	user = nil
// 	err = json.Unmarshal([]byte(`
// 		{
// 			"name":"Matheus Zabin",
// 			"email":"emailInvalido"
// 		}
// 	`), &user)
// 	require.ErrorIs(err, uservo.ErrEmailInvalidFormat)

// }
