package learn_golang_fiber

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRoutingGroup(t *testing.T) {

	getBooksHandler := func(ctx *fiber.Ctx) error {
		product := Product{
			Name:  "Books",
			Price: 100000,
		}

		return ctx.Status(fiber.StatusOK).JSON(product)
	}

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/books", getBooksHandler)

	v2 := api.Group("/v2")
	v2.Get("/books", getBooksHandler)

	request1 := httptest.NewRequest("GET", "/api/v1/books", nil)
	request1.Header.Set("Content-Type", "application/json")

	request2 := httptest.NewRequest("GET", "/api/v2/books", nil)
	request2.Header.Set("Content-Type", "application/json")

	response1, err1 := app.Test(request1)
	response2, err2 := app.Test(request2)

	assert.Nil(t, err1)
	assert.Nil(t, err2)

	body1, err1 := io.ReadAll(response1.Body)
	body2, err2 := io.ReadAll(response2.Body)

	assert.Nil(t, err1)
	assert.Nil(t, err2)

	var booksData1 = new(Product)
	var booksData2 = new(Product)

	json.Unmarshal(body1, booksData1)
	json.Unmarshal(body2, booksData2)

	assert.Equal(t, booksData1.Name, booksData2.Name)
}
