package database

import (
	"fmt"
	"log"

	"github.com/billowdev/golang-api-clinic-management/pkg/configs"
	"gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}

	var err error
	var dialector gorm.Dialector

	// Determine database driver from DB_URL
	// Assuming DB_URL format: "postgres://..." or "mysql://..."
	dsn := configs.DB_URL
	if dsn == "" {
		return nil, fmt.Errorf("DB_URL is not set")
	}

	// Simple check for database type
	if len(dsn) > 8 && dsn[:8] == "postgres" {
		dialector = postgres.Open(dsn)
	} else if len(dsn) > 5 && dsn[:5] == "mysql" {
		dialector = mysql.Open(dsn)
	} else {
		// Default to postgres if format is unclear
		dialector = postgres.Open(dsn)
	}

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	DB, err = gorm.Open(dialector, config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connection established successfully")
	return DB, nil
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database not initialized. Call InitDatabase() first.")
	}
	return DB
}

