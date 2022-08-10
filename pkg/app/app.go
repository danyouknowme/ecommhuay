package app

import "github.com/gofiber/fiber/v2"

func CreateFiberApp() *fiber.App {
	app := fiber.New()
	return app
}
