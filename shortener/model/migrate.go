package model

import "github.com/jinzhu/gorm"

func Migrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(
		&Link{},
	).Error
	return
}
