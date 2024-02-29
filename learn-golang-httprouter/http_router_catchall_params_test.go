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

func TestHttpRouterCatchAllParams(t *testing.T) {

	router := httprouter.New()

	router.GET("/products/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		productId := p.ByName("image")

		fmt.Fprint(w, productId)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/static/image/avatar.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, "/static/image/avatar.png", string(body))
}
