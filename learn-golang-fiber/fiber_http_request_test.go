package learn_golang_fiber

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHttpHeaderCookiesRequest(t *testing.T) {

	app.Get("/fullname", func(ctx *fiber.Ctx) error {
		firstName := ctx.Get("firstName", "John")
		lastName := ctx.Cookies("lastName", "Doe")

		return ctx.SendString(firstName + " " + lastName)
	})

	request := httptest.NewRequest("GET", "/fullname", nil)
	request.Header.Set("firstName", "Hutama")
	request.AddCookie(&http.Cookie{Name: "lastName", Value: "Trirahmanto"})

	response, err := app.Test(request)

	assert.Nil(t, err)

	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)

	assert.Equal(t, "Hutama Trirahmanto", string(body))
}
