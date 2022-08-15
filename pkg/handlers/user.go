package handlers

import (
	"log"

	"github.com/danyouknowme/ecommerce/pkg/database/dbmodels"
	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(ctx *fiber.Ctx) error {
	var req dbmodels.User
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	log.Printf("post: /api/v1/users/register")

	err := dbmodels.Register(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userResponse := dbmodels.UserResponse{
		Username: req.Username,
		FullName: req.FullName,
		Email:    req.Email,
	}

	return ctx.JSON(fiber.Map{
		"message": "Create a new user successfully!",
		"user":    userResponse,
	})
}

func LoginAPI(ctx *fiber.Ctx) error {
	var req dbmodels.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	log.Printf("post: /api/v1/users/login")
	user, err := dbmodels.Login(req.Username, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func GetUserAPI(ctx *fiber.Ctx) error {
	var req dbmodels.GetUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	log.Printf("get: /api/v1/users")

	user, err := dbmodels.GetUser(req.Username)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userResponse := dbmodels.UserResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	}

	return ctx.Status(fiber.StatusOK).JSON(userResponse)
}
