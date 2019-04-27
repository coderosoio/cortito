package connection

import (
	// Dialect for MySQL
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"common/config"
)

func getMySQLConnection(settings *config.Database) (db *gorm.DB, err error) {
	url := settings.URL(true, false)
	db, err = gorm.Open(string(settings.Adapter), url)
	return
}
