package connection

import (
	// Dialect for SQLite
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"common/config"
)

func getSQLiteConnection(settings *config.Database) (db *gorm.DB, err error) {
	path := settings.URL(false, false)
	db, err = gorm.Open(string(settings.Adapter), path)
	return
}
