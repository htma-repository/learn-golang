package learn_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FormPost(writer http.ResponseWriter, req *http.Request) {
	// err := req.ParseForm()

	// if err != nil {
	// 	panic(err)
	// }

	// firstName := req.PostForm.Get("first_name")
	// lastName := req.PostForm.Get("last_name")

	// if want use without parseForm, can use PostFormValue (simple way)
	firstName := req.PostFormValue("first_name")
	lastName := req.PostFormValue("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	reqBody := strings.NewReader("first_name=Hutama&last_name=Trirahmanto")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", reqBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	res := recorder.Result()
	b, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	assert.Equal(t, "Hello Hutama Trirahmanto", string(b))
}
