package main

import (
	"fmt"
	"gtest_example/app/internal/db"
	"gtest_example/app/utils/env"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	"github.com/wshaman/migrate"

	_ "gtest_example/app/internal/db/migrations"
)

func help() {
	c := os.Args[0]
	fmt.Printf(`Usage:
	%s command [params]
commands:
	help show this screen and exit
	up run all migrations 
	down rollback 1 last migration
	create creates a new migration file template 
Eg:
%s create add_table_users
`, c, c)
}

func dbConnect() (*db.DB, error) {
	port, _ := strconv.Atoi(env.EnvOrDef("DB_PORT", "5432"))
	c, err := db.NewWithDatabaseCreation("postgres",
		env.EnvOrDef("DB_USER", "postgres"),
		env.EnvOrDef("DB_PASSWD", "pass"),
		env.EnvOrDef("DB_HOST", "127.0.0.1"),
		port,
		env.EnvOrDef("DB_NAME", "tst"),
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(0)
	}
	dbObj, err := dbConnect()
	if err != nil {
		log.Fatal(err)
		//fmt.Println(err.Error())
	}
	command := strings.ToLower(os.Args[1])
	switch command {
	case "up":
		err = dbObj.MigrateUp()
	case "down":
		err = dbObj.MigrateDown()
	case "sync":
		err = dbObj.Sync()
	case "create":
		if len(os.Args) < 3 {
			help()
			os.Exit(1)
		}
		err = migrate.CreateFile(os.Args[2], "migrations", "./", true)
	default:
		help()
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}
