package learn_golang_fiber

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type Books struct {
	Name string `json:"name"`
}

func TestFiberMethod(t *testing.T) {
	app.Get("/books", func(c *fiber.Ctx) error {

		books := []Books{
			{
				Name: "Book-1",
			},
			{
				Name: "Book-2",
			},
		}

		return c.JSON(map[string]any{
			"status": "success",
			"data":   books,
		})
	})

	request := httptest.NewRequest("GET", "/books", nil)

	response, err := app.Test(request)

	assert.Nil(t, err)

	body, err := io.ReadAll(response.Body)

	var responseBody map[string]any

	app.Config().JSONDecoder(body, &responseBody)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "success", responseBody["status"])
	assert.Equal(t, "Book-1", responseBody["data"].([]any)[0].(map[string]any)["name"])
}
