package routes

import (
	"github.com/danyouknowme/ecommerce/pkg/handlers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func SetupRouter(app *fiber.App, secretKey string) {
	api := app.Group("/api/v1")
	authRoutes := app.Group("/api/auth/v1").Use(authRequired(secretKey))

	api.Get("/products", handlers.GetAllProductsAPI)
	api.Get("/products/:id", handlers.GetProductByIdAPI)
	api.Post("/products", handlers.AddNewProductAPI)
	api.Patch("/products/:id", handlers.UpdateProductAmountAPI)
	authRoutes.Get("/products/:id", handlers.DeleteProductByIdAPI)

	api.Post("/users/register", handlers.RegisterAPI)
	api.Post("/users/login", handlers.LoginAPI)
}

func authRequired(jwtSecret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: jwtError,
		SigningKey:   []byte(jwtSecret),
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
