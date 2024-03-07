package controller

import (
	"learn-golang-dependency-injection/helper"
	"learn-golang-dependency-injection/model/web"
	"learn-golang-dependency-injection/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller ProductControllerImpl) CreateProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCreateRequest := web.ProductCreateRequest{}
	helper.DecodeJSONFromRequest(request, &productCreateRequest)

	product := controller.ProductService.CreateProduct(request.Context(), productCreateRequest)

	productResponseJson := web.ResponseJson{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   product,
	}

	helper.EncodeJSONFromResponse(writer, productResponseJson)
}

func (controller ProductControllerImpl) FindAllProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	products := controller.ProductService.FindAllProduct(request.Context())

	productResponseJson := web.ResponseJson{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   products,
	}

	helper.EncodeJSONFromResponse(writer, productResponseJson)
}

func (controller ProductControllerImpl) FindByIdProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")

	id, err := strconv.Atoi(productId)
	helper.PanicError(err)

	product := controller.ProductService.FindProductById(request.Context(), id)

	productResponseJson := web.ResponseJson{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   product,
	}

	helper.EncodeJSONFromResponse(writer, productResponseJson)
}

func (controller ProductControllerImpl) UpdateProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productUpdateRequest := web.ProductUpdateRequest{}
	helper.DecodeJSONFromRequest(request, &productUpdateRequest)

	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicError(err)

	productUpdateRequest.Id = id

	updatedProduct := controller.ProductService.UpdateProduct(request.Context(), productUpdateRequest)

	productResponseJson := web.ResponseJson{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   updatedProduct,
	}

	helper.EncodeJSONFromResponse(writer, productResponseJson)
}

func (controller ProductControllerImpl) DeleteProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicError(err)

	controller.ProductService.DeleteProduct(request.Context(), id)

	productResponseJson := web.ResponseJson{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.EncodeJSONFromResponse(writer, productResponseJson)
}
