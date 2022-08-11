package handlers

import (
	"log"

	"github.com/danyouknowme/ecommerce/pkg/database/dbmodels"
	"github.com/gofiber/fiber/v2"
)

func GetAllProductsAPI(ctx *fiber.Ctx) error {
	log.Printf("get: /api/v1/products")

	products, err := dbmodels.GetAllProducts()
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(products)
}

func GetProductByIdAPI(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "Invalid product Id!",
		})
	}

	log.Printf("get: /api/v1/products/%d", id)

	product, err := dbmodels.GetProductById(id)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(product)
}

func AddNewProductAPI(ctx *fiber.Ctx) error {
	var req dbmodels.Product
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	log.Printf("post: /api/v1/products")

	err := dbmodels.AddNewProduct(req)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Add the new product successfully!",
	})
}

func UpdateProductAmountAPI(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "Invalid product Id!",
		})
	}

	log.Printf("patch: /api/v1/products/%d", id)

	err = dbmodels.UpdateProductAmount(id)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Update amount product successfully!",
	})
}
