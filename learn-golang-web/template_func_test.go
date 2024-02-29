package learn_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Users struct {
	Name    string
	Address string
}

func (u Users) SayGoodbye(name string) string {
	return "Hello " + u.Name + ", Your Address is " + u.Address + ". Goodbye " + name
}

func TemplateFunc(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayGoodbye "Budi"}}`))

	err := t.ExecuteTemplate(writer, "FUNCTION", Users{
		Name:    "Hutama",
		Address: "Tangerang",
	})

	if err != nil {
		panic(err)
	}

}

func TestTemplateFunc(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunc(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

func TemplateFuncGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))

	err := t.ExecuteTemplate(writer, "FUNCTION", Users{
		Name:    "Hutama",
		Address: "Tangerang",
	})

	if err != nil {
		panic(err)
	}

}

func TestTemplateFuncGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFuncGlobal(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

func TemplateFuncGlobalCreate(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")

	t.Funcs(map[string]any{
		"sum": func(num1, num2 int) int {
			return num1 + num2
		},
	}).Parse(`{{sum .Num1 .Num2}}`)

	err := t.ExecuteTemplate(writer, "FUNCTION", map[string]any{
		"Num1": 10,
		"Num2": 10,
	})

	if err != nil {
		panic(err)
	}
}

func TestTemplateFuncGlobalCreate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFuncGlobalCreate(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

func TemplateFuncGlobalCreatePipeline(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")

	t.Funcs(map[string]any{
		"sum": func(num1, num2 int) int {
			return num1 + num2
		},
		"mul": func(value int, mul int) int {
			return value * mul
		},
	}).Parse(`{{sum .Num1 .Num2 | mul .Num3}}`)

	err := t.ExecuteTemplate(writer, "FUNCTION", map[string]any{
		"Num1": 10,
		"Num2": 10,
		"Num3": 2,
	})

	if err != nil {
		panic(err)
	}
}

func TestTemplateFuncGlobalCreatePipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFuncGlobalCreatePipeline(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
