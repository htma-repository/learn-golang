package learn_golang_web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// go:embed templates/if_else.html
// var templateIf embed.FS

type DataIf struct {
	Name string
}

func TemplateActionIfElse(writer http.ResponseWriter, request *http.Request) {
	// t := template.Must(template.ParseFS(templateIf, "templates/if_else.html"))

	templateGlobal.ExecuteTemplate(writer, "if_else.html", DataIf{
		Name: "Hutama",
	})
}

func TestTemplateActionIfElse(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIfElse(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

//go:embed templates/comparison.html
var templateComparison embed.FS

func TemplateActionComparison(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templateComparison, "templates/comparison.html"))

	t.ExecuteTemplate(writer, "comparison.html", map[string]any{
		"Value": 2,
	})
}

func TestTemplateActionComparison(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionComparison(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

//go:embed templates/range.html
var templateRange embed.FS

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templateRange, "templates/range.html"))

	t.ExecuteTemplate(writer, "range.html", map[string]any{
		"Values": []string{
			"Hutama", "Trirahmanto", "Tama",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

//go:embed templates/with.html
var templateWith embed.FS

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templateWith, "templates/with.html"))

	t.ExecuteTemplate(writer, "with.html", map[string]any{
		"Address": map[string]any{
			"City":    "Tangerang",
			"Country": "Indonesia",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
