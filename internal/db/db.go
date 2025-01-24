package db

import (
	"log"

	"github.com/ratheeshkumar25/blog-crud-api/config"
	"github.com/ratheeshkumar25/blog-crud-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(config *config.Config) *gorm.DB {
	if config.Database_url == "" {
		log.Fatal("Database URL is not set in the configuration")
	}
	// Log the connection details
	log.Printf("Connecting to DB with URL: %s", config.Database_url)

	DB, err := gorm.Open(postgres.Open(config.Database_url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Printf("Error while connecting to the database: %v", err)
		return nil
	}

	// Migrate the schema
	if err := DB.AutoMigrate(&model.BlogPost{}); err != nil {
		log.Printf("Error while migrating: %v", err)
		return nil
	}

	return DB
}
