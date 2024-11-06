package Utils

import (
	"github.com/gofiber/fiber/v2"
)

func HasAccess(c *fiber.Ctx, aci string) bool {
	// acis := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["acis"]
	return true
}
