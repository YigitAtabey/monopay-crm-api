// controllers/userController.go
package controllers

import (
	"monopay-crm-api/config" // DB bağlantısını tuttuğun yapı
	"monopay-crm-api/models" // User model tanımı

	"github.com/gofiber/fiber/v2"
)

// GetUsers: Tüm kullanıcıların listesini dönen handler.
// - DB’den tüm User kayıtlarını çeker.
// - Hata varsa 500, yoksa JSON array döner.
func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	// DB.Find: tüm user kayıtlarını users değişkenine atar
	if err := config.DB.Find(&users).Error; err != nil {
		// DB okuma hatası
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kullanıcılar alınamadı",
		})
	}

	// Başarılı cevap: users dizisini JSON olarak dön
	return c.JSON(users)
}
