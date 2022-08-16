package handlers

import "github.com/gofiber/fiber/v2"

func WelcomeAPI(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to ecommhuay api!")
}
