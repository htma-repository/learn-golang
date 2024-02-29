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

func ResponseStatusCode(writer http.ResponseWriter, req *http.Request) {

	firstName := req.PostFormValue("first_name")

	if firstName != "" {
		fmt.Fprintf(writer, "first_name: %s", firstName)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "first_name: empty")
	}
}

func TestStatusCodeInvalid(t *testing.T) {
	reqBody := strings.NewReader("first_name=")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", reqBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	ResponseStatusCode(recorder, request)

	res := recorder.Result()
	b, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)

	assert.Equal(t, "first_name: empty", string(b))

}

func TestStatusCodeValid(t *testing.T) {
	reqBody := strings.NewReader("first_name=Hutama")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", reqBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	ResponseStatusCode(recorder, request)

	res := recorder.Result()
	b, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)

	assert.Equal(t, "first_name: Hutama", string(b))

}
