package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "Welcome to Ecommerce api"})
	})
}
