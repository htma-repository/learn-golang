package app

import (
	"learn-golang-restful-api/controller"
	"learn-golang-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/products", controller.CreateProduct)
	router.GET("/api/products", controller.FindAllProduct)
	router.GET("/api/products/:productId", controller.FindByIdProduct)
	router.PUT("/api/products/:productId", controller.UpdateProduct)
	router.DELETE("/api/products/:productId", controller.DeleteProduct)

	router.PanicHandler = exception.ErrorHandler

	return router
}
