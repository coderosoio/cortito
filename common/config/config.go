package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jinzhu/configor"
)

const (
	// Development environment
	Development Environment = "development"
	// Test environment
	Test Environment = "test"
	// Production environment
	Production Environment = "production"
)

var (
	// CurrentEnvironment of the application
	CurrentEnvironment = Development

	configurationFile = "config.yml"
	currentConfig     = (*Config)(nil)
)

// Environment represents an application environment.
type Environment string

// Config holds the current application configuration.
type Config struct {
	Debug    bool
	Database map[string]*Database
	Auth     *Auth
	Service  map[string]*Service
	Session  *Session
	HTTP     *HTTP
}

func init() {
	currentEnvironment := os.Getenv("APP_ENV")
	if currentEnvironment == "" {
		currentEnvironment = (string)(Development)
	}
	CurrentEnvironment = Environment(currentEnvironment)
}

// SetConfigurationFile sets the path to the configuration file to use.
func SetConfigurationFile(filename string) (err error) {
	if !filepath.IsAbs(filename) {
		if filename, err = filepath.Abs(filename); err != nil {
			return
		}
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			return fmt.Errorf("configuration file does not exist: %v", err)
		}
	}
	configurationFile = filename
	return
}

// GetConfig returns the current configuration loaded.
func GetConfig() (*Config, error) {
	if currentConfig == nil {
		var result Config
		err := configor.New(
			&configor.Config{
				ENVPrefix:   "APP",
				Environment: string(CurrentEnvironment),
				Debug:       !IsProduction(),
				Verbose:     !IsProduction(),
			},
		).Load(&result, configurationFile)
		if err != nil {
			return nil, fmt.Errorf("error loading configuration file %s: %v", configurationFile, err)
		}
		currentConfig = &result
	}
	return currentConfig, nil
}

// IsDevelopment checks if the current environment is `Development`.
func IsDevelopment() bool {
	return IsEnvironment(Development)
}

// IsTest checks if the current environment is `Test`.
func IsTest() bool {
	return IsEnvironment(Test)
}

// IsProduction checks if the current environment is `Production`.
func IsProduction() bool {
	return IsEnvironment(Production)
}

// IsEnvironment checks if the given `environment` is the same as the `CurrentEnvironment`.
func IsEnvironment(environment Environment) bool {
	return CurrentEnvironment == environment
}
