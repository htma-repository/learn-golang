package learn_golang_fiber

import (
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRouteParameter(t *testing.T) {

	app.Get("/first/:firstName/last/:lastName", func(ctx *fiber.Ctx) error {
		firstName := ctx.Params("firstName")
		lastName := ctx.Params("lastName")
		allParams := ctx.AllParams()

		fmt.Println(allParams["firstName"])
		fmt.Println(allParams["lastName"])

		return ctx.SendString(firstName + " " + lastName)
	})

	request := httptest.NewRequest("GET", "/first/Hutama/last/Trirahmanto", nil)

	response, err := app.Test(request)

	assert.Nil(t, err)

	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)

	assert.Equal(t, "Hutama Trirahmanto", string(body))
}
