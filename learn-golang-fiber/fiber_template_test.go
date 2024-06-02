package learn_golang_fiber

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
	"github.com/stretchr/testify/assert"
)

func TestTemplates(t *testing.T) {
	engine := mustache.New("./views", ".mustache")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/index", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "This is title",
			"body":  "This is body",
		})
	})

	request := httptest.NewRequest("GET", "/index", nil)

	response, err := app.Test(request)

	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)

	assert.Nil(t, err)

	assert.Contains(t, string(bytes), "This is title")
	assert.Contains(t, string(bytes), "This is body")
}
