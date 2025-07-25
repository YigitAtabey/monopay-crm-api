package controllers

import (
	"monopay-crm-api/config"
	"monopay-crm-api/models"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Kullanıcıyı ID ile engelle (role="blocked")
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

	if user.Role == "blocked" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Kullanıcı zaten engellenmiş"})
	}

	user.Role = "blocked"
	user.UpdatedAt = time.Now()
	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kullanıcı engellenemedi"})
	}

	return c.JSON(fiber.Map{"message": "Kullanıcı başarıyla engellendi"})
}

// Email ile admin yap (role="admin")
func MakeAdminByEmail(c *fiber.Ctx) error {
	var body struct {
		Email string `json:"email"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz istek formatı"})
	}
	email := strings.TrimSpace(body.Email)
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email parametresi gerekli"})
	}

	var user models.User
	if err := config.DB.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Kullanıcı bulunamadı"})
	}

	if user.Role == "admin" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Kullanıcı zaten admin"})
	}

	user.Role = "admin"
	user.UpdatedAt = time.Now()
	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kullanıcı admin yapılamadı"})
	}

	return c.JSON(fiber.Map{"message": "Kullanıcı başarıyla admin yapıldı"})
}

// Email ile engelle (role="blocked")
func BlockUserByEmail(c *fiber.Ctx) error {
	var body struct {
		Email string `json:"email"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz istek formatı"})
	}
	email := strings.TrimSpace(body.Email)
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email parametresi gerekli"})
	}

	var user models.User
	err := config.DB.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Kullanıcı bulunamadı"})
	}

	if user.Role == "blocked" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Kullanıcı zaten engellenmiş"})
	}

	user.Role = "blocked"
	user.UpdatedAt = time.Now()
	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kullanıcı engellenemedi"})
	}

	return c.JSON(fiber.Map{"message": "Kullanıcı başarıyla engellendi"})
}

// Email ile engeli kaldır (role="user")
func UnblockUserByEmail(c *fiber.Ctx) error {
	var body struct {
		Email string `json:"email"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz istek formatı"})
	}
	email := strings.TrimSpace(body.Email)
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email parametresi gerekli"})
	}

	var user models.User
	err := config.DB.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Kullanıcı bulunamadı"})
	}

	if user.Role != "blocked" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Kullanıcı engelli değil"})
	}

	user.Role = "user"
	user.UpdatedAt = time.Now()
	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kullanıcı engeli kaldırılamadı"})
	}

	return c.JSON(fiber.Map{"message": "Kullanıcı engeli kaldırıldı"})
}
