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

func GetProductByIdAPI(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "Invalid product Id!",
		})
	}

	log.Printf("/api/v1/products/%d", id)

	product, err := dbmodels.GetProductById(id)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(product)
}

func AddNewProductAPI(ctx *fiber.Ctx) error {
	var newProduct dbmodels.Product
	if err := ctx.BodyParser(&newProduct); err != nil {
		return err
	}

	err := dbmodels.AddNewProduct(newProduct)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Add the new product successfully!",
	})
}
