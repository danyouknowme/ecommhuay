package handlers

import (
	"github.com/danyouknowme/ecommerce/pkg/database/dbmodels"
	"github.com/gofiber/fiber/v2"
)

func AddOrUpdateProductInCartAPI(ctx *fiber.Ctx) error {
	err := dbmodels.AddOrUpdateProductInCart("dannyisadmin", 3, false)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "update the product in cart successfully!",
	})
}
