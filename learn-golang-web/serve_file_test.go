package learn_golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

//go:embed resources/found.html
var found string

//go:embed resources/notfound.html
var not_found string

func ServeFile(writer http.ResponseWriter, request *http.Request) {

	if request.URL.Query().Get("name") != "" {
		// with ServeFile
		// http.ServeFile(writer, request, "./resources/found.html")

		// with embed
		fmt.Fprint(writer, found)
	} else {
		// with ServeFile
		// http.ServeFile(writer, request, "./resources/notfound.html")

		// with embed
		fmt.Fprint(writer, not_found)
	}
}

func TestServeFile(t *testing.T) {
	// req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	// recorder := httptest.NewRecorder()

	// ServeFile(recorder, req)

	// res := recorder.Result()

	// b, err := io.ReadAll(res.Body)

	// if err != nil {
	// 	panic(err)
	// }

	// assert.NotEqual(t, "Hello World", string(b))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
