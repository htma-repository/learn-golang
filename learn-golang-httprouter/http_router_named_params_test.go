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

func TestHttpRouterNamedParams(t *testing.T) {

	router := httprouter.New()

	router.GET("/products/:productId/items/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		productId := p.ByName("productId")

		itemId := p.ByName("itemId")

		fmt.Fprintf(w, "%s %s", productId, itemId)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/7/items/10", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, "7 10", string(body))
}
