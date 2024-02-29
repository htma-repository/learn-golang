package learn_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "test data")
	})

	mux.HandleFunc("/data/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "test images")
	})

	server := http.Server{Addr: "localhost:8080", Handler: mux}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
