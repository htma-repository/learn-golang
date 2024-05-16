package learn_golang_fiber

import (
	"encoding/json"
	"errors"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestErrorHandling(t *testing.T) {

	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"message": err.Error(),
				})
			},
		},
	)

	app.Get("/error", func(ctx *fiber.Ctx) error {
		return errors.New("something went wrong")
	})

	request := httptest.NewRequest("GET", "/error", nil)

	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)

	var errorData fiber.Map

	json.Unmarshal(body, &errorData)

	assert.Equal(t, "error", errorData["status"])
	assert.Equal(t, "something went wrong", errorData["message"])
}
