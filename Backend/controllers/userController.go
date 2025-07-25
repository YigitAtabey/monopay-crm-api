package controllers

import (
	"monopay-crm-api/config"
	"monopay-crm-api/models"

	"github.com/gofiber/fiber/v2"
)

// GetUsers: Tüm kullanıcıların listesini döner.
// Sadece adminler tüm kullanıcıları görür, user sadece kendini görür, blocked hiçbir kullanıcıyı göremez.
func GetUsers(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(models.User)
	var users []models.User

	switch currentUser.Role {
	case "admin":
		// Admin ise tüm kullanıcıları getir
		if err := config.DB.Find(&users).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Kullanıcılar alınamadı",
			})
		}
		return c.JSON(users)
	case "user":
		// User ise sadece kendini görebilir
		if err := config.DB.Where("id = ?", currentUser.ID).Find(&users).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Kullanıcı alınamadı",
			})
		}
		return c.JSON(users)
	case "blocked":
		// Engelli kullanıcı hiçbir kullanıcıyı göremez
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Erişiminiz engellendi",
		})
	default:
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Geçersiz kullanıcı rolü",
		})
	}
}
