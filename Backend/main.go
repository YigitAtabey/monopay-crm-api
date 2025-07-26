package main

import (
	"log"
	"time"

	"monopay-crm-api/config"
	"monopay-crm-api/models"
	"monopay-crm-api/routes"
	"monopay-crm-api/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// 1. DB bağlantısını aç (yeni ConnectDB fonksiyonun döngüyle bağlantı deniyor)
	config.ConnectDB()

	// Ekstra garanti için: Bağlantı gerçekten var mı kontrolü (normalde gerek yok ama isterse eklenebilir)
	if config.DB == nil {
		log.Fatal("Veritabanı bağlantısı kurulamadı, uygulama başlatılamıyor!")
	}

	// 2. Gerekli tabloları oluştur
	if err := config.DB.AutoMigrate(&models.User{}, &models.Task{}, &models.TokenBlacklist{}); err != nil {
		log.Fatal("Migrate hatası:", err)
	}

	// 3. Default admin kullanıcısını oluştur (eğer yoksa)
	config.CreateDefaultAdmin()

	// 4. Fiber app başlat
	app := fiber.New(fiber.Config{
		// istersen buraya özel config ekleyebilirsin
	})

	// 5. CORS middleware: front-end'den gelen istekleri aç
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	// 6. Route’ları ekle
	routes.Setup(app)

	// 7. (Opsiyonel) Arka planda blacklist temizleme
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		for range ticker.C {
			services.CleanExpiredTokens()
		}
	}()

	// 8. Sunucuyu ayağa kaldır
	log.Fatal(app.Listen(":8083"))
}
