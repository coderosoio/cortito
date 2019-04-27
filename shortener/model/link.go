package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Link struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`
	URL       string `gorm:"not null"`
	Hash      string `gorm:"not null; unique_index"`
	Visits    int
	LastVisit *time.Time
}
