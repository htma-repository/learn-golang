package learn_golang_fiber

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRequestForm(t *testing.T) {

	app.Post("/signin", func(ctx *fiber.Ctx) error {
		email := ctx.FormValue("email")
		password := ctx.FormValue("password")

		return ctx.SendString("User sign-in with email " + email + ", and password is " + password)
	})

	requestBody := strings.NewReader("email=test@gmail.com&password=Test123$")
	request := httptest.NewRequest("POST", "/signin", requestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := app.Test(request)

	assert.Nil(t, err)

	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)

	assert.Equal(t, "User sign-in with email test@gmail.com, and password is Test123$", string(body))
}
