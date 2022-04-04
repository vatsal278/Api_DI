package config

import (
	"database/sql"
	"fmt"
)

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "goblog"
	connection_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", dbUser, dbPass, "localhost", "9001", dbName)

	db, err := sql.Open(dbDriver, connection_string)
	if err != nil {
		panic(err.Error())
	}
	return db
}
