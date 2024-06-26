// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"

	model "github.com/keshu12345/guardianlink/model"
)

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

// FetchUsers provides a mock function with given fields: c
func (_m *AuthService) FetchUsers(c *gin.Context) ([]model.User, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for FetchUsers")
	}

	var r0 []model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context) ([]model.User, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context) []model.User); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Singin provides a mock function with given fields: c
func (_m *AuthService) Singin(c *gin.Context) (string, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Singin")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context) (string, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context) string); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Singup provides a mock function with given fields: c
func (_m *AuthService) Singup(c *gin.Context) (string, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Singup")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context) (string, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context) string); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validate provides a mock function with given fields: c
func (_m *AuthService) Validate(c *gin.Context) {
	_m.Called(c)
}

// NewAuthService creates a new instance of AuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthService {
	mock := &AuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
