package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/wshaman/migrate"

	_ "gtest_example/app/internal/db/migrations"
)

type DB struct {
	Conn *sqlx.DB
}

var errDBDoesNotExists = errors.New("db does not exists")

func connString(dbType, user, pwd, dbHost string, port int, dbName string) string {
	if dbName == "" {
		return fmt.Sprintf("%s://%s:%s@%s:%d?sslmode=disable", dbType, user, pwd, dbHost, port)
	}
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable", dbType, user, pwd, dbHost, port, dbName)
}

func New(dbType, user, pwd, host string, port int, dbName string) (*DB, error) {
	c, err := sqlx.Connect(dbType, connString(dbType, user, pwd, host, port, dbName))
	if err != nil {
		if err.Error() == fmt.Sprintf("pq: database \"%s\" does not exist", dbName) {
			return nil, errDBDoesNotExists
		}
		return nil, errors.Wrap(err, "failed to connect to DB")
	}
	db := &DB{
		Conn: c,
	}
	return db, nil
}

func newWithDatabaseCreation(dbType, user, pwd, host string, port int, dbName string) (*DB, error) {
	c, err := sqlx.Connect(dbType, connString(dbType, user, pwd, host, port, ""))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to DB")
	}
	if _, err = c.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)); err != nil {
		return nil, errors.Wrap(err, "failed to create DB")
	}
	if err = c.Close(); err != nil {
		return nil, errors.Wrap(err, "failed to close DB connection")
	}
	return New(dbType, user, pwd, host, port, dbName)
}

func NewWithDatabaseCreation(dbType, user, pwd, host string, port int, dbName string) (*DB, error) {
	tmp, err := New(dbType, user, pwd, host, port, dbName)
	if err == errDBDoesNotExists {
		return newWithDatabaseCreation(dbType, user, pwd, host, port, dbName)
	}
	if err != nil {
		return nil, err
	}
	return tmp, nil
}

func (db *DB) MigrateUp() error {
	if err := migrate.Up(db.Conn.DB); err != nil {
		return errors.Wrap(err, "failed to migrate DB up")
	}
	return nil
}

func (db *DB) MigrateDown() error {
	if err := migrate.Down(db.Conn.DB); err != nil {
		return errors.Wrap(err, "failed to migrate DB up")
	}
	return nil
}

func (db *DB) Sync() error {
	if err := migrate.Sync(db.Conn.DB); err != nil {
		return errors.Wrap(err, "failed to sync DB up")
	}
	return nil
}
