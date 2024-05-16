package learn_golang_fiber

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	_ "embed"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

//go:embed example.txt
var fileEmbed []byte

func TestMultipartForm(t *testing.T) {

	app.Post("/upload", func(ctx *fiber.Ctx) error {
		file, err := ctx.FormFile("file")
		if err != nil {
			return err
		}

		err = ctx.SaveFile(file, "./assets/"+file.Filename)
		if err != nil {
			return err
		}

		return ctx.SendString("Upload success")
	})

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, err := writer.CreateFormFile("file", "example.txt")
	assert.Nil(t, err)
	file.Write(fileEmbed)
	writer.Close()

	request := httptest.NewRequest("POST", "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	response, err := app.Test(request)

	assert.Nil(t, err)

	data, err := io.ReadAll(response.Body)

	assert.Nil(t, err)

	assert.Equal(t, "Upload success", string(data))
}
