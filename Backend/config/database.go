package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"monopay-crm-api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var db *gorm.DB
	var err error

	for {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("❌ Veritabanına bağlanılamadı, tekrar denenecek:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		DB = db
		log.Println("✅ Veritabanı bağlantısı başarılı")
		break
	}
}

// CreateDefaultAdmin - Sistem başlatıldığında otomatik olarak default admin kullanıcısı oluşturur
func CreateDefaultAdmin() {
	var admin models.User
	// Eğer admin email'li kullanıcı zaten varsa, tekrar oluşturma
	if err := DB.Where("email = ?", "admin").First(&admin).Error; err == nil {
		log.Println("ℹ️  Default admin kullanıcısı zaten mevcut")
		return
	}

	// Admin şifresini hashle
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("1234"), 14)
	if err != nil {
		log.Printf("❌ Default admin şifresi hashlenemedi: %v", err)
		return
	}

	// Default admin kullanıcısını oluştur
	defaultAdmin := models.User{
		Name:     "Administrator",
		Email:    "admin",
		Password: string(hashedPassword),
		Role:     "admin",
	}

	if err := DB.Create(&defaultAdmin).Error; err != nil {
		log.Printf("❌ Default admin kullanıcısı oluşturulamadı: %v", err)
		return
	}

	log.Println("✅ Default admin kullanıcısı başarıyla oluşturuldu (email: admin, password: 1234)")
}
