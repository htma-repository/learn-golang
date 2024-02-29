package learn_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type ValidationMiddleware struct {
	Handler http.Handler
}

func (middleware *ValidationMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("Before middleware run")

	middleware.Handler.ServeHTTP(writer, request)

	fmt.Println("After middleware run")
}

type ErrorHandlerMiddleware struct {
	Handler http.Handler
}

func (middleware *ErrorHandlerMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	// error recover func
	defer func() {
		err := recover()
		fmt.Println("Recover: ", err)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}

	}()

	middleware.Handler.ServeHTTP(writer, request)
}

func TestMiddlewareServer(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware run")
		fmt.Fprint(w, "Hello Middleware")
	})

	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware run")
		panic("Error occurred")
	})

	validation := &ValidationMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandlerMiddleware{
		Handler: validation,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
