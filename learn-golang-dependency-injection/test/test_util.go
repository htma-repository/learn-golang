package test

import (
	"database/sql"
	"learn-golang-dependency-injection/app"
	"learn-golang-dependency-injection/dependency_injection"
	"learn-golang-dependency-injection/helper"
	"learn-golang-dependency-injection/middleware"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func SetupDbTest() *sql.DB {
	db, err := sql.Open("mysql", "hutamatr:Rahmanto123!@tcp(localhost:3306)/learn_golang_restful_api_test?parseTime=true")
	helper.PanicError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}

func TruncateDB(db *sql.DB) {
	_, err := db.Exec("TRUNCATE product")
	helper.PanicError(err)
}

func SetupRouter(db *sql.DB) http.Handler {
	validator := validator.New()

	controller := dependency_injection.InitializeController(db, validator)

	router := app.NewRouter(controller)

	return middleware.NewAuthMiddleware(router)
}
