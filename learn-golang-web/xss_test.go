package learn_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Go has default xss escape mechanism automatic if use html/template std lib

func TemplateXSS(writer http.ResponseWriter, request *http.Request) {
	err := templateGlobal.ExecuteTemplate(writer, "xss.html", map[string]any{
		"Title": "Template XSS",
		"Body":  "<p>This is xss demo</p>",
	})

	if err != nil {
		panic(err)
	}
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

func TemplateXSSOff(writer http.ResponseWriter, request *http.Request) {
	err := templateGlobal.ExecuteTemplate(writer, "xss.html", map[string]any{
		"Title": "Template XSS",
		"Body":  template.HTML("<p>This is xss demo</p>"),
	})

	if err != nil {
		panic(err)
	}
}

func TestTemplateXSSOff(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateXSSOff(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
