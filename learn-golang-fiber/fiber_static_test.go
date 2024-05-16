package learn_golang_fiber

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestStatic(t *testing.T) {
	app.Static("/static", "./assets/example.txt")

	request := httptest.NewRequest("GET", "/static", nil)

	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)

	assert.Equal(t, "This is test for fiber multipart file", string(body))
}
