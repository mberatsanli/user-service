package Controllers

import "github.com/gofiber/fiber/v2"

func Permissions(ctx *fiber.Ctx) error {
	return ctx.SendString("Permissions")
}
