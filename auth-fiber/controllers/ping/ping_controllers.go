package ping

import (
	"net/http"

	"github.com/gofiber/fiber"
)

// Ping controller to handle /fiber/ping
func Ping(c *fiber.Ctx) {
	c.Status(http.StatusOK)
	c.Send("Fiber Pong")
}
