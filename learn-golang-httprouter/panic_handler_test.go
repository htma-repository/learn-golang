package learn_golang_http_router

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	Method string
	Assert string
}

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprint(w, "Panic : ", i)
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic(map[string]any{
			"Status":  "error",
			"Message": "Error occurred",
		})
	})

	router.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic(map[string]any{
			"Status":  "error",
			"Message": "Failed post data",
		})
	})

	data := []TestConfig{
		{
			Method: http.MethodGet,
			Assert: "Panic : map[Message:Error occurred Status:error]",
		},
		{
			Method: http.MethodPost,
			Assert: "Panic : map[Message:Failed post data Status:error]",
		},
	}

	for _, test := range data {

		request := httptest.NewRequest(test.Method, "http://localhost:8080/", nil)
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)

		res := recorder.Result()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			panic(err)
		}

		assert.Equal(t, test.Assert, string(body))
	}

}
