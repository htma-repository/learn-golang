package main

import (
	"learn-golang-dependency-injection/helper"
	"learn-golang-dependency-injection/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddleware,
	}
}

func main() {

	server := InitializeServer()

	err := server.ListenAndServe()
	helper.PanicError(err)
}
