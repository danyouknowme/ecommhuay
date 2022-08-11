package routes

import (
	"github.com/danyouknowme/ecommerce/pkg/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/products", handlers.GetAllProductsAPI)
	api.Get("/products/:id", handlers.GetProductByIdAPI)
}
