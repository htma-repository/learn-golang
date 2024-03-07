//go:build wireinject
// +build wireinject

package main

import (
	"learn-golang-dependency-injection/app"
	"learn-golang-dependency-injection/controller"
	"learn-golang-dependency-injection/helper"
	"learn-golang-dependency-injection/middleware"
	"learn-golang-dependency-injection/repository"
	"learn-golang-dependency-injection/service"
	"net/http"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var productSet = wire.NewSet(
	repository.NewProductRepository,
	wire.Bind(new(repository.ProductRepository), new(*repository.ProductRepositoryImpl)),
	service.NewProductService,
	wire.Bind(new(service.ProductService), new(*service.ProductServiceImpl)),
	controller.NewProductController,
	wire.Bind(new(controller.ProductController), new(*controller.ProductControllerImpl)),
)

func InitializeServer() *http.Server {
	wire.Build(app.DB, helper.NewValidate, productSet, app.NewRouter, wire.Bind(new(http.Handler), new(*httprouter.Router)), middleware.NewAuthMiddleware, NewServer)
	return nil
}
