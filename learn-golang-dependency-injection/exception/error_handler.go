package exception

import (
	"learn-golang-dependency-injection/helper"
	"learn-golang-dependency-injection/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}
	if validationErr(writer, request, err) {
		return
	}
	internalServerErrorHandler(writer, request, err)

}

func validationErr(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {

	if exception, ok := err.(validator.ValidationErrors); ok {

		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		internalServerErrResponse := web.ResponseJson{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.EncodeJSONFromResponse(writer, internalServerErrResponse)

		return true
	}

	return false
}

func notFoundError(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {

	if exception, ok := err.(NotFoundError); ok {

		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		internalServerErrResponse := web.ResponseJson{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.EncodeJSONFromResponse(writer, internalServerErrResponse)

		return true
	}

	return false
}

func internalServerErrorHandler(writer http.ResponseWriter, _ *http.Request, err interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	internalServerErrResponse := web.ResponseJson{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.EncodeJSONFromResponse(writer, internalServerErrResponse)
}
