package models

import (
	"database/sql"
	"gtest_example/app/internal/db"

	"github.com/pkg/errors"
)

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"name"`
}

func rowsToUsers(rows *sql.Rows) (users []User, err error) {
	users = make([]User, 0)
	for rows.Next() {
		u := &User{}
		if err = rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, errors.Wrap(err, "failed to ListUsers (scan)")
		}
		users = append(users, *u)
	}
	return users, nil
}

func UserList(dbObj *db.DB) ([]User, error) {
	rows, err := dbObj.Conn.Query("select id, name, email from users")
	if err != nil {
		return nil, errors.Wrap(err, "failed to ListUsers")
	}
	return rowsToUsers(rows)
}

func UserListEmailLike(dbObj *db.DB, eml string) ([]User, error) {
	eml = eml + "%"
	rows, err := dbObj.Conn.Query("select id, name, email from users where email LIKE $1", eml)
	if err != nil {
		return nil, errors.Wrap(err, "failed to ListUsers")
	}
	return rowsToUsers(rows)
}

func insertUser(dbObj *db.DB, u *User) (err error) {
	var id int64
	q := "insert into users (name, email) values ($1, $2) returning id"
	if err = dbObj.Conn.QueryRow(q, u.Name, u.Email).Scan(&id); err != nil {
		return errors.Wrap(err, "failed to insert user")
	}
	u.ID = int(id)
	return nil
}

func updateUser(db *db.DB, u *User) error {
	q := "update users set  name=$1, email=$2 where id=$3;"
	if _, err := db.Conn.Exec(q, u.Name, u.Email, u.ID); err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}

func UserSave(db *db.DB, u *User) error {
	if u.ID == 0 {
		return insertUser(db, u)
	}
	return updateUser(db, u)
}
