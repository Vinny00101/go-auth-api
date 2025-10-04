package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name" binding:"required" gorm:"type:varchar(100);not null"`
	Email        string `json:"email" binding:"required,email" gorm:"type:varchar(100);not null;unique"`
	PasswordHash string `json:"password_hash" binding:"required,min=8" gorm:"type:varchar(255);not null"`
}
