package controllers

import (
	"log"
	"monopay-crm-api/config"
	"monopay-crm-api/models"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func BlockUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz kullanıcı ID"})
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Kullanıcı bulunamadı"})
	}

	user.IsBlocked = true
	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kullanıcı engellenemedi"})
	}

	return c.JSON(fiber.Map{"message": "Kullanıcı başarıyla engellendi"})
}

func MakeAdminByEmail(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email parametresi gerekli"})
	}

	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Kullanıcı bulunamadı"})
	}

	user.IsAdmin = true
	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kullanıcı admin yapılamadı"})
	}

	return c.JSON(fiber.Map{"message": "Kullanıcı başarıyla admin yapıldı"})
}

func BlockUserByEmail(c *fiber.Ctx) error {
	email := strings.TrimSpace(c.Query("email"))
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email parametresi gerekli"})
	}

	var user models.User
	err := config.DB.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Kullanıcı bulunamadı"})
	}

	err = config.DB.Model(&user).Updates(map[string]interface{}{
		"is_blocked": true,
		"updated_at": time.Now(),
	}).Error
	if err != nil {
		log.Println("Kullanıcı engelleme hatası:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kullanıcı engellenemedi"})
	}

	return c.JSON(fiber.Map{"message": "Kullanıcı başarıyla engellendi"})
}
