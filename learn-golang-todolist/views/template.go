package views

import (
	"embed"
	"net/http"
	"text/template"
)

//go:embed *.html
var templatesAll embed.FS

var TemplateGlobal = template.Must(template.ParseFS(templatesAll, "*.html"))

func TemplateForm(writer http.ResponseWriter, request *http.Request) {
	err := TemplateGlobal.ExecuteTemplate(writer, "form.html", nil)

	if err != nil {
		panic(err)
	}
}
