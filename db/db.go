package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"myapp/config"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", connectionString)

	if err != nil {
		panic("connectionString error..")
	}

	err = db.Ping()
	if err != nil {
		panic("DNS Invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
