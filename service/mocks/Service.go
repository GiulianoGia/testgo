// Code generated by mockery v2.33.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	types "gotest/types"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateGrocery provides a mock function with given fields: grocery
func (_m *Service) CreateGrocery(grocery *types.Grocery) (types.Grocery, error) {
	ret := _m.Called(grocery)

	var r0 types.Grocery
	var r1 error
	if rf, ok := ret.Get(0).(func(*types.Grocery) (types.Grocery, error)); ok {
		return rf(grocery)
	}
	if rf, ok := ret.Get(0).(func(*types.Grocery) types.Grocery); ok {
		r0 = rf(grocery)
	} else {
		r0 = ret.Get(0).(types.Grocery)
	}

	if rf, ok := ret.Get(1).(func(*types.Grocery) error); ok {
		r1 = rf(grocery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateGroceryForUser provides a mock function with given fields: username, groceryId
func (_m *Service) CreateGroceryForUser(username string, groceryId int) error {
	ret := _m.Called(username, groceryId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(username, groceryId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateNewUser provides a mock function with given fields: user
func (_m *Service) CreateNewUser(user types.User) types.User {
	ret := _m.Called(user)

	var r0 types.User
	if rf, ok := ret.Get(0).(func(types.User) types.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(types.User)
	}

	return r0
}

// DeleteGroceryById provides a mock function with given fields: id
func (_m *Service) DeleteGroceryById(id int) (types.Grocery, error) {
	ret := _m.Called(id)

	var r0 types.Grocery
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (types.Grocery, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) types.Grocery); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(types.Grocery)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGroceryForUser provides a mock function with given fields: username, groceryId
func (_m *Service) DeleteGroceryForUser(username string, groceryId int) error {
	ret := _m.Called(username, groceryId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(username, groceryId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUserByName provides a mock function with given fields: username
func (_m *Service) DeleteUserByName(username string) (types.User, error) {
	ret := _m.Called(username)

	var r0 types.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (types.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) types.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(types.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByUsernameAndPassword provides a mock function with given fields: username, password
func (_m *Service) FindUserByUsernameAndPassword(username string, password string) (types.User, error) {
	ret := _m.Called(username, password)

	var r0 types.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (types.User, error)); ok {
		return rf(username, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) types.User); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(types.User)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllGroceries provides a mock function with given fields:
func (_m *Service) GetAllGroceries() []types.Grocery {
	ret := _m.Called()

	var r0 []types.Grocery
	if rf, ok := ret.Get(0).(func() []types.Grocery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Grocery)
		}
	}

	return r0
}

// GetAllGroceriesFromUser provides a mock function with given fields: userId
func (_m *Service) GetAllGroceriesFromUser(userId string) ([]types.Grocery, error) {
	ret := _m.Called(userId)

	var r0 []types.Grocery
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]types.Grocery, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(string) []types.Grocery); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Grocery)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUsers provides a mock function with given fields:
func (_m *Service) GetAllUsers() []types.User {
	ret := _m.Called()

	var r0 []types.User
	if rf, ok := ret.Get(0).(func() []types.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.User)
		}
	}

	return r0
}

// GetGroceryByName provides a mock function with given fields: name
func (_m *Service) GetGroceryByName(name string) []types.Grocery {
	ret := _m.Called(name)

	var r0 []types.Grocery
	if rf, ok := ret.Get(0).(func(string) []types.Grocery); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Grocery)
		}
	}

	return r0
}

// GetRoleIdByName provides a mock function with given fields: username
func (_m *Service) GetRoleIdByName(username string) (int, error) {
	ret := _m.Called(username)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByName provides a mock function with given fields: name
func (_m *Service) GetUserByName(name string) (types.User, error) {
	ret := _m.Called(name)

	var r0 types.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (types.User, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) types.User); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(types.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserIdByUsername provides a mock function with given fields: username
func (_m *Service) GetUserIdByUsername(username string) (string, error) {
	ret := _m.Called(username)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchGroceriesFromUser provides a mock function with given fields: query
func (_m *Service) SearchGroceriesFromUser(query string) []types.Grocery {
	ret := _m.Called(query)

	var r0 []types.Grocery
	if rf, ok := ret.Get(0).(func(string) []types.Grocery); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Grocery)
		}
	}

	return r0
}

// UpdateGrocery provides a mock function with given fields: grocery
func (_m *Service) UpdateGrocery(grocery types.Grocery) (types.Grocery, error) {
	ret := _m.Called(grocery)

	var r0 types.Grocery
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Grocery) (types.Grocery, error)); ok {
		return rf(grocery)
	}
	if rf, ok := ret.Get(0).(func(types.Grocery) types.Grocery); ok {
		r0 = rf(grocery)
	} else {
		r0 = ret.Get(0).(types.Grocery)
	}

	if rf, ok := ret.Get(1).(func(types.Grocery) error); ok {
		r1 = rf(grocery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatusOfGrocery provides a mock function with given fields: groceryId, status
func (_m *Service) UpdateStatusOfGrocery(groceryId int, status bool) (types.Grocery, error) {
	ret := _m.Called(groceryId, status)

	var r0 types.Grocery
	var r1 error
	if rf, ok := ret.Get(0).(func(int, bool) (types.Grocery, error)); ok {
		return rf(groceryId, status)
	}
	if rf, ok := ret.Get(0).(func(int, bool) types.Grocery); ok {
		r0 = rf(groceryId, status)
	} else {
		r0 = ret.Get(0).(types.Grocery)
	}

	if rf, ok := ret.Get(1).(func(int, bool) error); ok {
		r1 = rf(groceryId, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: updatedUser
func (_m *Service) UpdateUser(updatedUser types.User) (types.User, error) {
	ret := _m.Called(updatedUser)

	var r0 types.User
	var r1 error
	if rf, ok := ret.Get(0).(func(types.User) (types.User, error)); ok {
		return rf(updatedUser)
	}
	if rf, ok := ret.Get(0).(func(types.User) types.User); ok {
		r0 = rf(updatedUser)
	} else {
		r0 = ret.Get(0).(types.User)
	}

	if rf, ok := ret.Get(1).(func(types.User) error); ok {
		r1 = rf(updatedUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
