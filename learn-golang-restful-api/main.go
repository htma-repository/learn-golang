package main

import (
	"learn-golang-restful-api/app"
	"learn-golang-restful-api/controller"
	"learn-golang-restful-api/helper"
	"learn-golang-restful-api/middleware"
	"learn-golang-restful-api/repository"
	"learn-golang-restful-api/service"
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
