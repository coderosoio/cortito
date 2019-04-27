package connection

import (
	"github.com/jinzhu/gorm"

	// Dialect for PostgreSQL
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"common/config"
)

func getPostgreSQLConnection(settings *config.Database) (db *gorm.DB, err error) {
	url := settings.URL(true, false)
	db, err = gorm.Open(string(settings.Adapter), url)
	return
}
