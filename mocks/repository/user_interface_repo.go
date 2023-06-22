// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	entity "crud/entity"

	mock "github.com/stretchr/testify/mock"
)

// UserInterfaceRepo is an autogenerated mock type for the UserInterfaceRepo type
type UserInterfaceRepo struct {
	mock.Mock
}

type UserInterfaceRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *UserInterfaceRepo) EXPECT() *UserInterfaceRepo_Expecter {
	return &UserInterfaceRepo_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: user
func (_m *UserInterfaceRepo) CreateUser(user *entity.User) (*entity.User, error) {
	ret := _m.Called(user)

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.User) (*entity.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(*entity.User) *entity.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserInterfaceRepo_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type UserInterfaceRepo_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - user *entity.User
func (_e *UserInterfaceRepo_Expecter) CreateUser(user interface{}) *UserInterfaceRepo_CreateUser_Call {
	return &UserInterfaceRepo_CreateUser_Call{Call: _e.mock.On("CreateUser", user)}
}

func (_c *UserInterfaceRepo_CreateUser_Call) Run(run func(user *entity.User)) *UserInterfaceRepo_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.User))
	})
	return _c
}

func (_c *UserInterfaceRepo_CreateUser_Call) Return(_a0 *entity.User, _a1 error) *UserInterfaceRepo_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserInterfaceRepo_CreateUser_Call) RunAndReturn(run func(*entity.User) (*entity.User, error)) *UserInterfaceRepo_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUser provides a mock function with given fields: email
func (_m *UserInterfaceRepo) DeleteUser(email string) (interface{}, error) {
	ret := _m.Called(email)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (interface{}, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserInterfaceRepo_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type UserInterfaceRepo_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - email string
func (_e *UserInterfaceRepo_Expecter) DeleteUser(email interface{}) *UserInterfaceRepo_DeleteUser_Call {
	return &UserInterfaceRepo_DeleteUser_Call{Call: _e.mock.On("DeleteUser", email)}
}

func (_c *UserInterfaceRepo_DeleteUser_Call) Run(run func(email string)) *UserInterfaceRepo_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UserInterfaceRepo_DeleteUser_Call) Return(_a0 interface{}, _a1 error) *UserInterfaceRepo_DeleteUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserInterfaceRepo_DeleteUser_Call) RunAndReturn(run func(string) (interface{}, error)) *UserInterfaceRepo_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserById provides a mock function with given fields: id
func (_m *UserInterfaceRepo) GetUserById(id uint) (entity.User, error) {
	ret := _m.Called(id)

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (entity.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) entity.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserInterfaceRepo_GetUserById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserById'
type UserInterfaceRepo_GetUserById_Call struct {
	*mock.Call
}

// GetUserById is a helper method to define mock.On call
//   - id uint
func (_e *UserInterfaceRepo_Expecter) GetUserById(id interface{}) *UserInterfaceRepo_GetUserById_Call {
	return &UserInterfaceRepo_GetUserById_Call{Call: _e.mock.On("GetUserById", id)}
}

func (_c *UserInterfaceRepo_GetUserById_Call) Run(run func(id uint)) *UserInterfaceRepo_GetUserById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *UserInterfaceRepo_GetUserById_Call) Return(_a0 entity.User, _a1 error) *UserInterfaceRepo_GetUserById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserInterfaceRepo_GetUserById_Call) RunAndReturn(run func(uint) (entity.User, error)) *UserInterfaceRepo_GetUserById_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: user, id
func (_m *UserInterfaceRepo) UpdateUser(user *entity.User, id uint) (interface{}, error) {
	ret := _m.Called(user, id)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.User, uint) (interface{}, error)); ok {
		return rf(user, id)
	}
	if rf, ok := ret.Get(0).(func(*entity.User, uint) interface{}); ok {
		r0 = rf(user, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.User, uint) error); ok {
		r1 = rf(user, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserInterfaceRepo_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type UserInterfaceRepo_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - user *entity.User
//   - id uint
func (_e *UserInterfaceRepo_Expecter) UpdateUser(user interface{}, id interface{}) *UserInterfaceRepo_UpdateUser_Call {
	return &UserInterfaceRepo_UpdateUser_Call{Call: _e.mock.On("UpdateUser", user, id)}
}

func (_c *UserInterfaceRepo_UpdateUser_Call) Run(run func(user *entity.User, id uint)) *UserInterfaceRepo_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.User), args[1].(uint))
	})
	return _c
}

func (_c *UserInterfaceRepo_UpdateUser_Call) Return(_a0 interface{}, _a1 error) *UserInterfaceRepo_UpdateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserInterfaceRepo_UpdateUser_Call) RunAndReturn(run func(*entity.User, uint) (interface{}, error)) *UserInterfaceRepo_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserInterfaceRepo creates a new instance of UserInterfaceRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserInterfaceRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserInterfaceRepo {
	mock := &UserInterfaceRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}