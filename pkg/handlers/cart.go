package handlers

import (
	"github.com/danyouknowme/ecommhuay/pkg/database/dbmodels"
	"github.com/gofiber/fiber/v2"
)

func AddOrUpdateProductInCartAPI(ctx *fiber.Ctx) error {
	var req dbmodels.ProductInCartRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := dbmodels.AddOrUpdateProductInCart(req.Username, req.ProductId, req.IsAddedQuantity)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "update the product in cart successfully!",
	})
}
