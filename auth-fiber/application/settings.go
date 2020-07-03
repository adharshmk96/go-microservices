package application

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func loadSettings() {
	app.Settings.ErrorHandler = errorHandler
}

func errorHandler(ctx *fiber.Ctx, err error) {
	fmt.Println(err)

	ctx.JSON(err)
}
