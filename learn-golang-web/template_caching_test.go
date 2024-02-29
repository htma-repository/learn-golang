package learn_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.html
var templatesAll embed.FS

var templateGlobal = template.Must(template.ParseFS(templatesAll, "templates/*.html"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	err := templateGlobal.ExecuteTemplate(writer, "index.html", "This is template caching")

	if err != nil {
		panic(err)
	}
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
