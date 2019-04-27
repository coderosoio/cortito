package connection

import (
	"fmt"

	"github.com/jinzhu/gorm"

	cfg "common/config"
)

func GetDatabaseConnection(name string) (db *gorm.DB, err error) {
	config, err := cfg.GetConfig()
	if err != nil {
		return nil, err
	}
	settings, found := config.Database[name]
	if !found {
		return nil, fmt.Errorf("no database conneciton found for %s", name)
	}
	var timezoneCommand string
	switch settings.Adapter {
	case cfg.MySQLDatabaseAdapter:
		timezoneCommand = "SET time_zone = '+00:00'"
		db, err = getMySQLConnection(settings)
	case cfg.PostgreSQLDatabaseAdapter:
		timezoneCommand = "SET timezone = 'UTC'"
		db, err = getPostgreSQLConnection(settings)
	case cfg.SQLiteDatabaseAdapter:
		db, err = getSQLiteConnection(settings)
	default:
		return nil, fmt.Errorf("database adapter %s not supported for connection %s", settings.Adapter, name)
	}
	if err != nil {
		return
	}
	if config.Debug {
		db.Debug()
	}
	db.LogMode(config.Debug)
	db.SingularTable(settings.Singularize)
	if settings.UTC {
		if _, err = db.DB().Exec(timezoneCommand); err != nil {
			return nil, fmt.Errorf("error setting UTC timezone: %v", err)
		}
	}
	return
}
