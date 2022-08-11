package handlers

import (
	"log"

	"github.com/danyouknowme/ecommerce/pkg/database/dbmodels"
	"github.com/gofiber/fiber/v2"
)

func GetAllProductsAPI(ctx *fiber.Ctx) error {
	log.Println("/api/v1/products")
	products, err := dbmodels.GetAllProducts()
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(products)
}
