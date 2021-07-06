package models

import (
	"gtest_example/app/internal/db"
)

type userMock struct {
}

func (um userMock) UserList(_ *db.DB) ([]User, error) {
	return []User{
		{
			ID:    0,
			Name:  "stub",
			Email: "stub@example.com",
		},
	}, nil
}

func (um userMock) UserListEmailLike(_ *db.DB, _ string) ([]User, error) {
	return []User{
		{
			ID:    0,
			Name:  "stub",
			Email: "stub@example.com",
		},
	}, nil
}

func (um userMock) UserSave(_ *db.DB, _ *User) error {
	return nil
}
