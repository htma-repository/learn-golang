package learn_golang_fiber

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func TestHttpResponse(t *testing.T) {
	app.Get("/product", func(ctx *fiber.Ctx) error {
		product := Product{
			Name:  "Books",
			Price: 100000,
		}

		return ctx.Status(fiber.StatusOK).JSON(product)
	})

	request := httptest.NewRequest("GET", "/product", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)

	var productData = new(Product)

	err = json.Unmarshal(body, productData)
	assert.Nil(t, err)

	assert.Equal(t, "Books", productData.Name)
	assert.Equal(t, float64(100000), productData.Price)
}
