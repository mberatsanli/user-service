package Middleware

import (
	"github.com/gofiber/fiber/v2"
	"user/app/Utils"
)

func AccessMiddleware(aci string) fiber.Handler {
	successHandler := func(ctx *fiber.Ctx) error {
		hasUsersViewAccess := Utils.HasAccess(ctx, aci)
		if !hasUsersViewAccess {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "access-error",
				"message": "You do not have access to this resource! You need (" + aci + ") access.",
			})
		}

		return ctx.Next()
	}

	return AuthMiddleware(successHandler)
}
