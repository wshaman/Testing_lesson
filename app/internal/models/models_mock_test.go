// +build mock

package models

import (
	db2 "gtest_example/app/internal/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	var db *db2.DB
	ul, err := Query.UserList(db)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(ul))
	assert.Equal(t, "stub@example.com", ul[0].Email)
}
