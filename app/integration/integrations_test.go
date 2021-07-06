// +build integration

package integration

import (
	"fmt"
	"gtest_example/app/integration/testdata"
	"gtest_example/app/internal/models"
	"gtest_example/app/internal/user"
	"gtest_example/app/utils/waitfor"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/pkg/errors"

	_ "github.com/lib/pq"

	"gtest_example/app/internal/db"
	"gtest_example/app/utils/env"
)

type integTest struct {
	DbName     string
	DbUser     string
	DbPassword string
	DbPort     int
	DbHost     string
	Db         *db.DB
}

var it integTest

func (i *integTest) dbInit() (err error) {
	if i.Db, err = db.NewWithDatabaseCreation("postgres",
		i.DbUser,
		i.DbPassword,
		i.DbHost,
		i.DbPort,
		i.DbName,
	); err != nil {
		return err
	}
	if err = i.Db.MigrateUp(); err != nil {
		return errors.Wrap(err, "failed to migrate db")
	}
	return nil
}

func (i *integTest) dbDrop() error {
	q := fmt.Sprintf("drop database %s", i.DbName)
	if _, err := i.Db.Conn.Exec(q); err != nil {
		return errors.Wrap(err, "failed to drop DB")
	}
	return nil
}

func (i *integTest) dbClose() error {
	if i.Db == nil {
		return nil
	}
	if err := i.Db.Conn.Close(); err != nil {
		return errors.Wrap(err, "can't close connection")
	}
	return nil
}

func (i *integTest) dbReset() (err error) {
	if err = i.dbClose(); err != nil {
		return err
	}
	if i.Db, err = db.New("postgres",
		i.DbUser,
		i.DbPassword,
		i.DbHost,
		i.DbPort,
		"",
	); err != nil {
		return err
	}

	if err = i.dbDrop(); err != nil {
		return err
	}
	return i.dbInit()
}

func (i *integTest) setDbEnvs() (err error) {
	if i.DbPort, err = strconv.Atoi(env.EnvOrDef("DB_PORT", "5432")); err != nil {
		return errors.Wrap(err, "failed to set db creds")
	}
	i.DbName = env.EnvOrDef("DB_NAME", "tst")
	i.DbHost = env.EnvOrDef("DB_HOST", "127.0.0.1")
	i.DbUser = env.EnvOrDef("DB_USER", "postgres")
	i.DbPassword = env.EnvOrDef("DB_PASSWORD", "pass")
	return nil
}

func testInit(m *testing.M) int {
	if err := it.setDbEnvs(); err != nil {
		log.Fatal(err)
	}
	if err := waitfor.It(it.DbHost, it.DbPort, 10); err != nil {
		log.Fatal(err)
	}
	if err := it.dbInit(); err != nil {
		log.Fatal(err)
	}
	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(testInit(m))
}

func TestOne(t *testing.T) {
	err := it.dbReset()
	require.NoError(t, err)
	for _, v := range testdata.CrewData {
		u, err := user.Register(it.Db, v.InName)
		require.NoError(t, err)
		assert.Equal(t, v.OutEmail, u.Email)
	}
	users, err := models.Query.UserList(it.Db)
	require.NoError(t, err)
	assert.Equal(t, len(testdata.CrewData), len(users))
}

func TestTwo(t *testing.T) {
	err := it.dbReset()
	require.NoError(t, err)
	for _, v := range testdata.SameNameData {
		_, err := user.Register(it.Db, v.InName)
		require.NoError(t, err)
	}
	users, err := models.Query.UserList(it.Db)
	require.NoError(t, err)
	assert.Equal(t, len(testdata.CrewData), len(users))
}
