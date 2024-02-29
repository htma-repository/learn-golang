package learn_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HelloHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello Http")
}

// Create valid unit test for http test, this unit test no need to run server, recommended if want to run http unit test
func TestHttpTest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/test", nil)
	recorder := httptest.NewRecorder()

	HelloHTTP(recorder, req)

	response := recorder.Result()

	result, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	resultString := string(result)

	assert.Equal(t, "Hello Http", resultString)
	assert.NotNil(t, resultString)

}
