package models

import (
	"os"
	"strings"
)

type QueryCaller interface {
	UserModel
}

type q struct {
	UserModel
}

var Query QueryCaller

func init() {
	Query = q{userModel{}}
	if useMock := os.Getenv("USE_MOCK"); strings.ToLower(useMock) == "true" {
		Query = q{userMock{}}
	}
}
