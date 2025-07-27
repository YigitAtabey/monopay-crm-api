package config

import (
	"fmt"
	"log"
	"os"
	"time"

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

// CreateDefaultAdmin varsayılan admin kullanıcısını oluşturur
func CreateDefaultAdmin() {
	// Önce admin kullanıcısının var olup olmadığını kontrol et
	var adminUser struct {
		ID    uint   `gorm:"primarykey"`
		Email string `gorm:"unique"`
		Role  string
	}

	// Admin email ile kullanıcı var mı kontrol et
	err := DB.Select("id, email, role").Where("email = ?", "admin").First(&adminUser).Error
	if err == nil {
		// Admin zaten var
		log.Printf("ℹ️  Varsayılan admin kullanıcısı zaten mevcut (ID: %d, Role: %s)", adminUser.ID, adminUser.Role)
		return
	}

	// Admin kullanıcı yok, oluştur
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("1234"), 14)
	if err != nil {
		log.Printf("❌ Admin şifresi hashlenemedi: %v", err)
		return
	}

	// Admin kullanıcısını oluştur - models.User import etmemek için inline struct kullanıyoruz
	adminData := map[string]interface{}{
		"name":       "Administrator",
		"email":      "admin",
		"password":   string(hashedPassword),
		"role":       "admin",
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	if err := DB.Table("users").Create(&adminData).Error; err != nil {
		log.Printf("❌ Varsayılan admin kullanıcısı oluşturulamadı: %v", err)
		return
	}

	log.Println("✅ Varsayılan admin kullanıcısı oluşturuldu (Email: admin, Şifre: 1234)")
}
