package config

import (
	"fmt"
	"os"

	// Auto load env variables. Needed in local development
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

// Config contains Application configuration data, and other shared resources
type Config struct {
	DBName string
	DBHost string
	DBUser string
	DBPass string
	Logger *zap.Logger
}

// Get returns a new Application configuration.
// It is encouraged to call this function once and then pass the value as parameter
func Get() *Config {
	logger, _ := zap.NewProduction()
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	return &Config{
		DBName: dbName,
		DBHost: dbHost,
		DBUser: dbUser,
		DBPass: dbPass,
		Logger: logger,
	}
}

// GetDBConnStr returns a database connection string
func (c *Config) GetDBConnStr() string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", c.DBHost, c.DBUser, c.DBPass, c.DBName)
}

// GetLogger returns the configuration logger
func (c *Config) GetLogger() *zap.Logger {
	return c.Logger
}
