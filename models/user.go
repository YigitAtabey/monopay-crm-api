package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"name" validate:"required"`
	Email     string `gorm:"unique" json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	IsBlocked bool   `gorm:"default:false"`
	IsAdmin   bool   `json:"is_admin" gorm:"default:false"`
}
