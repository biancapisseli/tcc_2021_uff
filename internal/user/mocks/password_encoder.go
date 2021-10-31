// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	uservo "ifoodish-store/internal/user/domain/valueobject"
)

// PasswordEncoder is an autogenerated mock type for the PasswordEncoder type
type PasswordEncoder struct {
	mock.Mock
}

// EncodePassword provides a mock function with given fields: rawPassword
func (_m *PasswordEncoder) EncodePassword(rawPassword uservo.PasswordRaw) (uservo.PasswordEncoded, error) {
	ret := _m.Called(rawPassword)

	var r0 uservo.PasswordEncoded
	if rf, ok := ret.Get(0).(func(uservo.PasswordRaw) uservo.PasswordEncoded); ok {
		r0 = rf(rawPassword)
	} else {
		r0 = ret.Get(0).(uservo.PasswordEncoded)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uservo.PasswordRaw) error); ok {
		r1 = rf(rawPassword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}