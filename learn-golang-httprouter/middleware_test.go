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

type ValidationMiddleware struct {
	http.Handler
}

func (middleware *ValidationMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("middleware running")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("middleware running")
}

func TestMiddleware(t *testing.T) {

	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("handler running")
		fmt.Fprint(w, "Middleware")
	})

	validationMiddleware := &ValidationMiddleware{router}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	validationMiddleware.ServeHTTP(recorder, request)

	res := recorder.Result()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Middleware", string(body))
}
