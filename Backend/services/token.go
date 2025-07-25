package services

import (
	"log"
	"monopay-crm-api/config"
	"monopay-crm-api/models"
	"time"
)

// CleanExpiredTokens: Süresi dolmuş tokenları blacklist tablosundan siler.
func CleanExpiredTokens() {
	log.Println("CleanExpiredTokens çalıştı")
	result := config.DB.Where("expires_at < ?", time.Now()).Delete(&models.TokenBlacklist{})
	if result.Error != nil {
		log.Println("Blacklist temizleme hatası:", result.Error)
	} else {
		log.Printf("Blacklist temizlendi, silinen kayıt sayısı: %d\n", result.RowsAffected)
	}
}
