package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController interface {
	CreateProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByIdProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
