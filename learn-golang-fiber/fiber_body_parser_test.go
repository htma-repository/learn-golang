package learn_golang_fiber

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type RegisterRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Email    string `json:"email" xml:"email" form:"email"`
	Password string `json:"password" xml:"password" form:"password"`
}

func TestRequestBodyParser(t *testing.T) {

	app.Post("/signup", func(ctx *fiber.Ctx) error {
		registerData := RegisterRequest{}

		err := ctx.BodyParser(&registerData)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("Failed Register")
		}

		return ctx.SendString("User sign-up with username " + registerData.Username + ", email " + registerData.Email + ", and password is " + registerData.Password)
	})

	t.Run("JSON", func(t *testing.T) {
		requestBody := strings.NewReader(`{"username":"testing","email": "test@example.com",       "password": "Test123$"}`)
		request := httptest.NewRequest("POST", "/signup", requestBody)
		request.Header.Set("Content-Type", "application/json")

		response, err := app.Test(request)

		assert.Nil(t, err)

		body, err := io.ReadAll(response.Body)

		assert.Nil(t, err)

		assert.Equal(t, "User sign-up with username testing, email test@example.com, and password is Test123$", string(body))
	})

	t.Run("FORM", func(t *testing.T) {
		requestBody := strings.NewReader("username=testing&email=test@example.com&password=Test123$")
		request := httptest.NewRequest("POST", "/signup", requestBody)
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		response, err := app.Test(request)

		assert.Nil(t, err)

		body, err := io.ReadAll(response.Body)

		assert.Nil(t, err)

		assert.Equal(t, "User sign-up with username testing, email test@example.com, and password is Test123$", string(body))
	})

	t.Run("XML", func(t *testing.T) {
		requestBody := strings.NewReader(`
		<RegisterRequest>
		<username>testing</username>
		<email>test@example.com</email>
		<password>Test123$</password>
		</RegisterRequest>
		`)
		request := httptest.NewRequest("POST", "/signup", requestBody)
		request.Header.Set("Content-Type", "application/xml")

		response, err := app.Test(request)

		assert.Nil(t, err)

		body, err := io.ReadAll(response.Body)

		assert.Nil(t, err)

		assert.Equal(t, "User sign-up with username testing, email test@example.com, and password is Test123$", string(body))
	})

}
