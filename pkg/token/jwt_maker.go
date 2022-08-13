package token

import (
	"github.com/danyouknowme/ecommerce/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func CreateToken(username string) (string, error) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		return "", err
	}

	payload, err := NewPayload(username)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(config.TokenSymmetricKey))
}

func AuthRequired() fiber.Handler {
	config, err := util.LoadConfig("../..")
	if err != nil {
		return nil
	}

	return jwtware.New(jwtware.Config{
		ErrorHandler:   AuthError,
		SigningKey:     []byte(config.TokenSymmetricKey),
		SuccessHandler: AuthSuccess,
	})
}

func AuthSuccess(ctx *fiber.Ctx) error {
	return ctx.Next()
}

func AuthError(ctx *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return ctx.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
