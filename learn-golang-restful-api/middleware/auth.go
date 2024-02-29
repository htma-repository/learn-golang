package middleware

import (
	"learn-golang-restful-api/helper"
	"learn-golang-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-Key") == "AUTHSECRET" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		internalServerErrResponse := web.ResponseJson{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.EncodeJSONFromResponse(writer, internalServerErrResponse)
	}
}
