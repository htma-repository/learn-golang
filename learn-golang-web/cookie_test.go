package learn_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, req *http.Request) {
	cookie := new(http.Cookie)

	cookie.Name = "X-HTMA-Name"
	cookie.Value = "Hutama Trirahmanto"
	cookie.Path = "/"

	http.SetCookie(writer, cookie)

	fmt.Fprint(writer, "Success add cookie")
}

func GetCookie(writer http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("X-HTMA-Name")

	if err != nil {
		fmt.Fprint(writer, "No cookie")
	} else {
		fmt.Fprintf(writer, "Cookie is %s", cookie.Value)
	}

}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, c := range cookies {
		fmt.Printf("%s : %s\n", c.Name, c.Value)
	}

}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-HTMA-Name"
	cookie.Value = "Hutama Trirahmanto"
	cookie.Path = "/"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	res := recorder.Result()

	b, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

}
