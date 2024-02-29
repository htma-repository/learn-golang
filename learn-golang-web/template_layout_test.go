package learn_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// go:embed templates/*.html
// var templateLayout embed.FS

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	// t := template.Must(template.ParseFS(templateLayout, "templates/head.html", "templates/layout.html", "templates/footer.html"))

	templateGlobal.ExecuteTemplate(writer, "layout", map[string]any{
		"Title": "Template Layout",
		"Name":  "Hutama Trirahmanto",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
