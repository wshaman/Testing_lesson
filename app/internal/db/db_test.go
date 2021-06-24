package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_connString(t *testing.T) {
	expected := "postgres://user:pass@localhost:5432/test_db?sslmode=disable"
	got := connString("postgres", "user", "pass", "localhost", 5432, "test_db")
	assert.Equal(t, expected, got)
}

func Test_connString_nodb(t *testing.T) {
	expected := "postgres://user:pass@localhost:5432?sslmode=disable"
	got := connString("postgres", "user", "pass", "localhost", 5432, "")
	assert.Equal(t, expected, got)
}
