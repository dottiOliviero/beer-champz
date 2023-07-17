package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDb :
func GetDb() (*gorm.DB, error) {
	dsn := "host=localhost user=test password=test dbname=test port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
