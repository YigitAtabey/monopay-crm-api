package middleware

import (
	"monopay-crm-api/models"

	"github.com/gofiber/fiber/v2"
)

func RequireAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User) // RequireAuth zaten user’ı local’a koyuyor

	if !user.IsAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Admin yetkiniz yok"})
	}

	return c.Next()
}
