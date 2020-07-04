package application

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

var (
	app = fiber.New()
)

// StartApplication start the fiber binding the routes
// and listens to 3000 port
func StartApplication() {
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	userRoutes(app)

	app.Listen(3000)
}
