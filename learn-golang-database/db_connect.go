package learn_golang_database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Rahmanto123!@tcp(localhost:3306)/learn_golang_database?parseTime=true")

	if err != nil {
		fmt.Println(errors.New(err.Error()))
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}
