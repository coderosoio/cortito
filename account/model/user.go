package model

import (
	"github.com/jinzhu/gorm"
)

// User model.
type User struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Email        string `gorm:"not null; unique_index"`
	PasswordHash string `gorm:"not null; unique_index"`
}
