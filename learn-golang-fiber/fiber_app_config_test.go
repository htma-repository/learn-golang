package learn_golang_fiber

import (
	"fmt"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestFiberConfig(t *testing.T) {
	app := fiber.New(fiber.Config{
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		AppName:      "Learn Fiber App",
		Prefork:      true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("success")
	})

	if fiber.IsChild() {
		fmt.Println("child process")
	} else {
		fmt.Println("parent process")
	}

	err := app.Listen(":8080")
	assert.Nil(t, err)

	assert.NotNil(t, app)
	assert.Equal(t, 10*time.Second, app.Config().IdleTimeout)
	assert.Equal(t, 10*time.Second, app.Config().ReadTimeout)
	assert.Equal(t, 10*time.Second, app.Config().WriteTimeout)
	assert.Equal(t, "Learn Fiber App", app.Config().AppName)
}
