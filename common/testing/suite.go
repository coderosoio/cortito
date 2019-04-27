package testing

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	// `require` fails tests immediately
	_ "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	commonConfig "common/config"
	"common/connection"
)

type Suite struct {
	suite.Suite

	Config *commonConfig.Config
	wd     string
}

func (suite *Suite) Init() {
	r := suite.Require()

	wd, err := os.Getwd()
	r.Nil(err)

	parent := filepath.Dir(wd)
	configFile := strings.Join([]string{parent, "config.yml"}, string(os.PathSeparator))

	err = commonConfig.SetConfigurationFile(configFile)
	r.Nil(err)

	config, err := commonConfig.GetConfig()
	r.Nil(err)

	suite.Config = config

	suite.wd = wd
}

func (suite *Suite) DeleteDatabaseFile(connectionName string) {
	r := suite.Require()

	settings, found := suite.Config.Database[connectionName]
	r.Truef(found, "database settings not found for %s", connectionName)

	databaseFile := strings.Join([]string{suite.wd, settings.Database}, string(os.PathSeparator))
	if _, err := os.Stat(databaseFile); os.IsNotExist(err) {
		return
	}
	err := os.Remove(databaseFile)
	r.Nilf(err, "error removing database file: %v", err)
}

func (suite *Suite) TruncateTables(connectionName string, tableNames ...string) {
	r := suite.Require()

	settings, found := suite.Config.Database[connectionName]
	r.Truef(found, "connection settings not found for %s", connectionName)

	db, err := connection.GetDatabaseConnection(connectionName)
	r.Nilf(err, "error getting database connection for %s: %v", connectionName, err)

	sql := "TRUNCATE TABLE %s"
	if settings.Adapter == commonConfig.SQLiteDatabaseAdapter {
		sql = "DELETE FROM %s"
	}
	for _, tableName := range tableNames {
		statement := fmt.Sprintf(sql, tableName)
		_, err := db.DB().Exec(statement)
		r.Nilf(err, "error truncating table %s: %v", tableName, err)
	}
}

func (suite *Suite) MigrateModels(connectionName string, models ...interface{}) {
	r := suite.Require()

	db, err := connection.GetDatabaseConnection(connectionName)
	r.Nilf(err, "error getting database connection for %s: %v", connectionName, err)

	err = db.AutoMigrate(models...).Error
	r.Nilf(err, "error migrating models: %v", err)
}
