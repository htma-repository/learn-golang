package learn_golang_fiber

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCtx(t *testing.T) {
	app.Get("/ctx", func(ctx *fiber.Ctx) error {
		query := ctx.Query("name", "Guest")

		return ctx.SendString("Hello " + query)
	})

	request := httptest.NewRequest("GET", "/ctx?name=Hutama", nil)

	response, err := app.Test(request)

	assert.Nil(t, err)

	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)

	assert.Equal(t, "Hello Hutama", string(body))

	// Guest

	request = httptest.NewRequest("GET", "/ctx", nil)

	response, err = app.Test(request)

	assert.Nil(t, err)

	body, err = io.ReadAll(response.Body)

	assert.Nil(t, err)

	assert.Equal(t, "Hello Guest", string(body))
}
