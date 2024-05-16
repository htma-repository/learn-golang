package learn_golang_fiber

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	app.Get("/file", func(ctx *fiber.Ctx) error {
		return ctx.Download("./assets/example.txt", "example.txt")
	})

	request := httptest.NewRequest("GET", "/file", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "attachment; filename=\"example.txt\"", response.Header.Get("Content-Disposition"))

	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)

	assert.Equal(t, "This is test for fiber multipart file", string(body))
}
