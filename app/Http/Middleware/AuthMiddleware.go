package Middleware

import (
	"encoding/base64"
	"fmt"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"user/config"
)

func AuthMiddleware(successHandler ...fiber.Handler) fiber.Handler {
	var handler = func(ctx *fiber.Ctx) error {
		return ctx.Next()
	}
	if len(successHandler) > 0 {
		handler = successHandler[0]
	}

	secret := []byte(config.Config("JWT_SECRET"))
	encodedKey := base64.StdEncoding.EncodeToString(secret)

	decodedKey, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		fmt.Println("Secret key decoding error:", err)
	}

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: "HS256",
			Key:    decodedKey,
		},
		ErrorHandler: jwtError,
		SuccessHandler: func(ctx *fiber.Ctx) error {
			return handler(ctx)
		},
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}

	c.Status(fiber.StatusUnauthorized)
	return c.JSON(fiber.Map{
		"status":  "error",
		"message": "Invalid or expired JWT",
		"data":    err.Error(),
	})
}
