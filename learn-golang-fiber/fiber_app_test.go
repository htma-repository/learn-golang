package learn_golang_fiber

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var app = fiber.New()

func TestFiberApp(t *testing.T) {
	assert.NotNil(t, app)
}
