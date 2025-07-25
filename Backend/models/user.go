package models

import (
	"gorm.io/gorm"
)

// Kullanıcı modeli: Sadece Role ile yönetim, eski IsBlocked ve IsAdmin kaldırıldı.
type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Email    string `gorm:"unique" json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role"` // "admin", "user", "blocked"
}
