package main

import (
	"learn-golang-migrate/app"
	"learn-golang-migrate/controller"
	"learn-golang-migrate/helper"
	"learn-golang-migrate/middleware"
	"learn-golang-migrate/repository"
	"learn-golang-migrate/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.DB()

	validator := validator.New()

	repository := repository.NewProductRepository()
	service := service.NewProductService(repository, db, validator)
	controller := controller.NewProductController(service)

	router := app.NewRouter(controller)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicError(err)
}
