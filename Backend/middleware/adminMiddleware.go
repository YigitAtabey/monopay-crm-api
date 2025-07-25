package middleware

import (
	"monopay-crm-api/models"

	"github.com/gofiber/fiber/v2"
)

// Admin gerektiren işlemler için: Sadece role="admin" geçer, blocked/user erişemez.
func RequireAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	if user.Role == "blocked" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Erişiminiz engellendi"})
	}
	if user.Role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Admin yetkiniz yok"})
	}

	return c.Next()
}

// Normal kullanıcı veya admin gerektiren işlemler için (blocked erişemez).
func RequireUserOrAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	if user.Role == "blocked" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Erişiminiz engellendi"})
	}
	// admin ve user erişebilir
	return c.Next()
}
