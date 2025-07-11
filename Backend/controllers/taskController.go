package controllers

import (
	"monopay-crm-api/config"
	"monopay-crm-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User) // JWT'den gelen kullanıcıyı alıyoruz

	var tasks []models.Task
	if err := config.DB.Where("user_id = ?", user.ID).Find(&tasks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Görevler alınamadı",
		})
	}

	return c.JSON(tasks)
}
func CreateTask(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

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

func UpdateTask(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	taskID := c.Params("id")

	var task models.Task
	// Önce görevi kullanıcıya ait ve ID eşleşmeli olarak bul
	if err := config.DB.Where("id = ? AND user_id = ?", taskID, user.ID).First(&task).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Görev bulunamadı"})
	}

	// Güncellenecek veriyi al
	var input struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz veri"})
	}

	// Değişiklik varsa güncelle
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

func DeleteTask(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	taskID := c.Params("id")

	var task models.Task
	// Kullanıcıya ait görevi bul
	if err := config.DB.Where("id = ? AND user_id = ?", taskID, user.ID).First(&task).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Görev bulunamadı"})
	}

	if err := config.DB.Delete(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Görev silinemedi"})
	}

	return c.JSON(fiber.Map{"message": "Görev silindi"})
}
