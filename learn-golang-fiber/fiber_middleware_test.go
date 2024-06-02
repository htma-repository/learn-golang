package learn_golang_fiber

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	middleware := func(c *fiber.Ctx) error {
		fmt.Printf("middleware running\n")
		id := c.Query("id")
		fmt.Println(id)
		return c.Next()
	}

	app.Get("/product", middleware, func(ctx *fiber.Ctx) error {
		product := Product{
			Name:  "Books",
			Price: 100000,
		}

		fmt.Println("handler running")

		return ctx.Status(fiber.StatusOK).JSON(product)
	})

	request := httptest.NewRequest("GET", "/product?id=1", nil)
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
