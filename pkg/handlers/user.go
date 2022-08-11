package handlers

import (
	"log"

	"github.com/danyouknowme/ecommerce/pkg/database/dbmodels"
	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(ctx *fiber.Ctx) error {
	var newUser dbmodels.User
	if err := ctx.BodyParser(&newUser); err != nil {
		return err
	}

	log.Printf("post: /api/v1/users/register")

	err := dbmodels.Register(newUser)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Create a new user successfully!",
	})
}
