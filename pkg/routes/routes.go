package routes

import (
	"github.com/danyouknowme/ecommhuay/pkg/handlers"
	"github.com/danyouknowme/ecommhuay/pkg/token"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, secretKey string) {
	app.Get("/", handlers.WelcomeAPI)

	api := app.Group("/api/v1")

	api.Get("/products", handlers.GetAllProductsAPI)
	api.Get("/products/:id", handlers.GetProductByIdAPI)
	api.Post("/products", handlers.AddNewProductAPI)
	api.Patch("/products/:id", handlers.UpdateProductAmountAPI)

	api.Post("/users/register", handlers.RegisterAPI)
	api.Post("/users/login", handlers.LoginAPI)

	api.Use(token.AuthRequired())

	api.Delete("/products/:id", handlers.DeleteProductByIdAPI)
	api.Get("/users/:username", handlers.GetUserAPI)
	api.Post("/carts", handlers.AddOrUpdateProductInCartAPI)
}
