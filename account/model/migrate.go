package model

import (
	"github.com/jinzhu/gorm"
)

// Migrate all models.
func Migrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(
		&User{},
		&UserToken{},
	).Error
	return
}
