package model

import "github.com/jinzhu/gorm"

type UserToken struct {
	gorm.Model
	Token  string `gorm:"not null; unique_index"`
	UserID uint   `gorm:"not null"`
	User   *User
}
