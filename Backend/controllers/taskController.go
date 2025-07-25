package controllers

import (
	"monopay-crm-api/config"
	"monopay-crm-api/models"

	"github.com/gofiber/fiber/v2"
)

// Sadece admin tüm görevleri görebilir, normal kullanıcı sadece kendi görevlerini görebilir
func GetTasks(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	var tasks []models.Task
	var err error

	if user.Role == "admin" {
		// Admin ise tüm görevleri getir
		err = config.DB.Find(&tasks).Error
	} else {
		// Engelli kullanıcı hiçbir şey göremez
		if user.Role == "blocked" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Erişiminiz engellendi"})
		}
		// User ise sadece kendi görevlerini getir
		err = config.DB.Where("user_id = ?", user.ID).Find(&tasks).Error
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Görevler alınamadı",
		})
	}

	return c.JSON(tasks)
}

// Engelli kullanıcı görev oluşturamaz, admin ve user oluşturabilir
func CreateTask(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	if user.Role == "blocked" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Erişiminiz engellendi"})
	}

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "İstek verisi hatalı"})
	}

	task.UserID = user.ID

	if err := config.DB.Create(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Görev oluşturulamadı"})
	}

	return c.JSON(task)
}

// Admin istediği görevi güncelleyebilir, user sadece kendi görevini
func UpdateTask(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	taskID := c.Params("id")

	var task models.Task

	if user.Role == "admin" {
		// Admin her görevi güncelleyebilir
		if err := config.DB.Where("id = ?", taskID).First(&task).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Görev bulunamadı"})
		}
	} else {
		if user.Role == "blocked" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Erişiminiz engellendi"})
		}
		// User ise sadece kendi görevini
		if err := config.DB.Where("id = ? AND user_id = ?", taskID, user.ID).First(&task).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Görev bulunamadı"})
		}
	}

	var input struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz veri"})
	}

	if input.Title != nil {
		task.Title = *input.Title
	}
	if input.Description != nil {
		task.Description = *input.Description
	}

	if err := config.DB.Save(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Görev güncellenemedi"})
	}

	return c.JSON(task)
}

// Admin istediği görevi silebilir, user sadece kendi görevini
func DeleteTask(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	taskID := c.Params("id")

	var task models.Task

	if user.Role == "admin" {
		// Admin her görevi silebilir
		if err := config.DB.Where("id = ?", taskID).First(&task).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Görev bulunamadı"})
		}
	} else {
		if user.Role == "blocked" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Erişiminiz engellendi"})
		}
		// User ise sadece kendi görevini
		if err := config.DB.Where("id = ? AND user_id = ?", taskID, user.ID).First(&task).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Görev bulunamadı"})
		}
	}

	if err := config.DB.Delete(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Görev silinemedi"})
	}

	return c.JSON(fiber.Map{"message": "Görev silindi"})
}
