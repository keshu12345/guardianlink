// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"

	model "github.com/keshu12345/guardianlink/model"
)

// NodeAService is an autogenerated mock type for the NodeAService type
type NodeAService struct {
	mock.Mock
}

// Create provides a mock function with given fields: c
func (_m *NodeAService) Create(c *gin.Context) (model.Block, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 model.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context) (model.Block, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context) model.Block); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(model.Block)
	}

	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: c
func (_m *NodeAService) Fetch(c *gin.Context) ([]model.Block, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []model.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context) ([]model.Block, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context) []model.Block); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c
func (_m *NodeAService) Update(c *gin.Context) (model.Block, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 model.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context) (model.Block, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context) model.Block); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(model.Block)
	}

	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewNodeAService creates a new instance of NodeAService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNodeAService(t interface {
	mock.TestingT
	Cleanup(func())
}) *NodeAService {
	mock := &NodeAService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
