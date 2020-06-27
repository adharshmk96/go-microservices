package main

import (
	"github.com/gofiber/fiber"
)

func authRouteHandle(c *fiber.Ctx) {
	c.Send("Hello what is cooking ?")
}

func main() {
	app := fiber.New()

	app.Get("/auth", authRouteHandle)

	app.Listen(3000)
}
