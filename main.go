package main

import (
	"log"
	"monopay-crm-api/config"
	"monopay-crm-api/models"
	"monopay-crm-api/routes"
	"monopay-crm-api/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()

	// Otomatik tablo oluştur
	if err := config.DB.AutoMigrate(&models.User{}, &models.Task{}, &models.TokenBlacklist{}); err != nil {
		log.Fatal("Migrate hatası:", err)
	}

	app := fiber.New()

	// Route'ları ekle
	routes.Setup(app)

	// Arkaplanda blacklist temizleme işlemini her saniye çalıştır
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		for {
			<-ticker.C
			services.CleanExpiredTokens()
		}
	}()

	// Sunucuyu başlat
	log.Fatal(app.Listen(":8083"))
}
