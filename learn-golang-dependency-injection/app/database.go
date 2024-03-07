package app

import (
	"database/sql"
	"learn-golang-dependency-injection/helper"

	"time"
)

func DB() *sql.DB {
	db, err := sql.Open("mysql", "hutamatr:Rahmanto123!@tcp(localhost:3306)/learn_golang_restful_api?parseTime=true")
	helper.PanicError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}
