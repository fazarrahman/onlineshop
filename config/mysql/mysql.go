package mysql

import (
	"log"

	"github.com/fazarrahman/onlineshop/lib"

	"github.com/jmoiron/sqlx"
)

// New ...
func New() (*sqlx.DB, error) {
	var password string = lib.GetEnv("DB_PASSWORD")
	var server string = lib.GetEnv("DB_SERVER")
	var dbName string = lib.GetEnv("DB_DATABASE_NAME")
	var username string = lib.GetEnv("DB_USERNAME")

	db, err := sqlx.Connect("mysql", username+":"+password+"@("+server+")/"+dbName+"?parseTime=true")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return db, nil
}
