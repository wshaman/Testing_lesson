// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	db "gtest_example/app/internal/db"

	mock "github.com/stretchr/testify/mock"

	models "gtest_example/app/internal/models"
)

// QueryCaller is an autogenerated mock type for the QueryCaller type
type QueryCaller struct {
	mock.Mock
}

// UserList provides a mock function with given fields: _a0
func (_m *QueryCaller) UserList(_a0 *db.DB) ([]models.User, error) {
	ret := _m.Called(_a0)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(*db.DB) []models.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*db.DB) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserListEmailLike provides a mock function with given fields: dbObj, eml
func (_m *QueryCaller) UserListEmailLike(dbObj *db.DB, eml string) ([]models.User, error) {
	ret := _m.Called(dbObj, eml)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(*db.DB, string) []models.User); ok {
		r0 = rf(dbObj, eml)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*db.DB, string) error); ok {
		r1 = rf(dbObj, eml)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserSave provides a mock function with given fields: _a0, u
func (_m *QueryCaller) UserSave(_a0 *db.DB, u *models.User) error {
	ret := _m.Called(_a0, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(*db.DB, *models.User) error); ok {
		r0 = rf(_a0, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
