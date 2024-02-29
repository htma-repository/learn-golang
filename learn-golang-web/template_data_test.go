package learn_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// go:embed templates/template_data.html
// var templateData embed.FS

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	// t := template.Must(template.ParseFS(templateData, "templates/template_data.html"))

	templateGlobal.ExecuteTemplate(writer, "template_data.html", map[string]any{
		"Title":       "Template data map",
		"Name":        "Hutama Trirahmanto",
		"Description": "This is template data map",
		"Address": map[string]any{
			"City": "Tangerang",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

type Address struct {
	City string
}

type Data struct {
	Title, Name, Description string
	Address                  Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	// t := template.Must(template.ParseFS(templateData, "templates/template_data.html"))

	templateGlobal.ExecuteTemplate(writer, "template_data.html", Data{
		Title:       "Template data struct",
		Name:        "Hutama Trirahmanto",
		Description: "This is template data map",
		Address: Address{
			City: "Tangerang",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
