package main

import (
	"log"
	"monopay-crm-api/config"
	"monopay-crm-api/models"
	"monopay-crm-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()

	// Otomatik tablo oluştur
	if err := config.DB.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatal("Migrate hatası:", err)
	}

	app := fiber.New()

	// Route'ları ekle
	routes.Setup(app)

	// Sunucuyu başlat
	log.Fatal(app.Listen(":8083"))
}
