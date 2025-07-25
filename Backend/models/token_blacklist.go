package models

import (
	"time"
)

// TokenBlacklist modeli token'ın kara listeye alınıp alınmadığını tutar.
// Role tabanlı sistemde burada değişiklik gerekmiyor.
type TokenBlacklist struct {
	ID            uint      `gorm:"primaryKey"`
	Token         string    `gorm:"type:varchar(512);unique"`
	BlacklistedAt time.Time `gorm:"autoCreateTime"`
	ExpiresAt     time.Time `gorm:"index"`
}

func (TokenBlacklist) TableName() string {
	return "token_blacklist"
}
