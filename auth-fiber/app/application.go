package app

import (
	"github.com/gofiber/fiber"
)

var (
	router = fiber.New()
)

// StartApplication start the app binding the routes
// and listens to 3000 port
func StartApplication() {
	mapUrls()
	router.Listen(3000)
}
