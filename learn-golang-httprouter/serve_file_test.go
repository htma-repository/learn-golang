package learn_golang_http_router

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var serveFile embed.FS

func TestHttpRouterServeFile(t *testing.T) {
	router := httprouter.New()

	dir, err := fs.Sub(serveFile, "resources")

	if err != nil {
		panic(err)
	}

	router.ServeFiles("/file/*filepath", http.FS(dir))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/file/serveFile.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, "ServeFile\n", string(body))
}
