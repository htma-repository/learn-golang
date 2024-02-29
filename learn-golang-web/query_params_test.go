package learn_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func QueryParams(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query().Get("name")
	fmt.Fprint(w, params)
}

func TestQueryParams(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "http:localhost:8080?name=hutama", nil)
	recorder := httptest.NewRecorder()

	QueryParams(recorder, req)

	res := recorder.Result()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	paramString := string(data)

	fmt.Println(paramString)

	assert.Equal(t, "hutama", paramString)
	assert.NotEqual(t, "trirahmanto", paramString)

}

func MultipleQueryParams(w http.ResponseWriter, req *http.Request) {
	firstName := req.URL.Query().Get("first_name")
	lastName := req.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http:localhost:8080?first_name=hutama&last_name=trirahmanto", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParams(recorder, req)

	res := recorder.Result()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	paramString := string(data)

	fmt.Println(paramString)

	assert.Equal(t, "Hello hutama trirahmanto", paramString)
	assert.NotEqual(t, "Hello hutama", paramString)
}

func MultipleValueQueryParams(w http.ResponseWriter, req *http.Request) {
	value := req.URL.Query()
	fmt.Println(value)
	params := value["name"]
	fmt.Println(params)

	fmt.Fprintf(w, "name-1 %s, name-2 %s", params[0], params[1])
}

func TestMultipleValueQueryParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=hutama&name=trirahmanto", nil)
	recorder := httptest.NewRecorder()

	MultipleValueQueryParams(recorder, req)

	res := recorder.Result()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	paramString := string(data)

	fmt.Println(paramString)

	assert.Equal(t, "name-1 hutama, name-2 trirahmanto", paramString)
	assert.NotEqual(t, "Hello hutama", paramString)
}
