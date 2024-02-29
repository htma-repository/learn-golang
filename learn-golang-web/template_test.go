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

func Template(writer http.ResponseWriter, request *http.Request) {

	textTemplate := "<html><body>{{.}}</body></html>"

	t := template.Must(template.New("New").Parse(textTemplate))

	t.ExecuteTemplate(writer, "New", "This is template")
}

func TestTemplateText(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	Template(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

func TemplateFile(writer http.ResponseWriter, request *http.Request) {

	// t := template.Must(template.ParseFiles("./templates/index.html"))

	templateGlobal.ExecuteTemplate(writer, "index.html", "This is template file")
}

func TestTemplateFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFile(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {

	// t := template.Must(template.ParseGlob("./templates/*.html"))

	templateGlobal.ExecuteTemplate(writer, "index.html", "This is template file")
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

//go:embed templates/*.html
var templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {

	// t := template.Must(template.ParseFS(templates, "templates/*.html"))

	templateGlobal.ExecuteTemplate(writer, "index.html", "This is template file")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
