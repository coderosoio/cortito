package config

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thoas/go-funk"
)

const (
	// PostgreSQLDatabaseAdapter is the adapter name for PostgreSQL.
	PostgreSQLDatabaseAdapter DatabaseAdapter = "postgres"
	// MySQLDatabaseAdapter is the adapter name for MySQL.
	MySQLDatabaseAdapter DatabaseAdapter = "mysql"
	// SQLiteDatabaseAdapter is the adapter name for SQLite.
	SQLiteDatabaseAdapter DatabaseAdapter = "sqlite3"
)

// DatabaseAdapter represents a database adapter constant.
type DatabaseAdapter string

// Database holds database connection parameters.
type Database struct {
	Adapter     DatabaseAdapter `default:"postgres"`
	Hostname    string
	Port        int
	Username    string
	Password    string
	Database    string
	Singularize bool   `default:"false"`
	UTC         bool   `default:"true"`
	Description string `default:""`
	Params      map[string]interface{}
}

// URL returns a connection string for the database.
func (d *Database) URL(withAdapter bool, wrapHost bool) string {
	var buffer bytes.Buffer
	if d.Adapter == SQLiteDatabaseAdapter {
		return d.Database
	}
	if withAdapter {
		buffer.WriteString(fmt.Sprintf("%s://", d.Adapter))
	}
	buffer.WriteString(d.Username)
	if len(d.Password) > 0 {
		buffer.WriteString(fmt.Sprintf(":%s", d.Password))
	}
	hostAndPort := d.Hostname
	if d.Port > 0 {
		hostAndPort = fmt.Sprintf("%s:%d", d.Hostname, d.Port)
	}
	if wrapHost {
		buffer.WriteString(fmt.Sprintf("@tcp(%s)", hostAndPort))
	} else {
		buffer.WriteString(fmt.Sprintf("@%s", hostAndPort))
	}
	buffer.WriteString(fmt.Sprintf("/%s", d.Database))
	if len(d.Params) > 0 {
		queryString := strings.Join(funk.Map(d.Params, func(name string, value interface{}) string {
			return fmt.Sprintf("%s=%s", name, value)
		}).([]string), "&")
		buffer.WriteString(fmt.Sprintf("?%s", queryString))
	}
	return buffer.String()
}
