package learn_golang_fiber

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TestRequestBody(t *testing.T) {

	app.Post("/signin", func(ctx *fiber.Ctx) error {
		body := ctx.Body()
		var signInData = SignIn{}
		err := json.Unmarshal(body, &signInData)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("Failed to signin")
		}

		return ctx.SendString("User sign-in with email " + signInData.Email + ", and password is " + signInData.Password)
	})

	requestBody := strings.NewReader(`{"email": "test@example.com", "password": "Test123$"}`)
	request := httptest.NewRequest("POST", "/signin", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)

	assert.Nil(t, err)

	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)

	assert.Equal(t, "User sign-in with email test@example.com, and password is Test123$", string(body))
}
