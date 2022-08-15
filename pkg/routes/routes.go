package routes

import (
	"github.com/danyouknowme/ecommerce/pkg/handlers"
	"github.com/danyouknowme/ecommerce/pkg/token"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, secretKey string) {
	api := app.Group("/api/v1")

	api.Get("/products", handlers.GetAllProductsAPI)
	api.Get("/products/:id", handlers.GetProductByIdAPI)
	api.Post("/products", handlers.AddNewProductAPI)
	api.Patch("/products/:id", handlers.UpdateProductAmountAPI)

	api.Post("/users/register", handlers.RegisterAPI)
	api.Post("/users/login", handlers.LoginAPI)

	api.Use(token.AuthRequired())

	api.Delete("/products/:id", handlers.DeleteProductByIdAPI)
	api.Get("/users", handlers.GetUserAPI)
	api.Post("/carts", handlers.AddOrUpdateProductInCartAPI)
}
