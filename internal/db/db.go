package db

import (
	"log"

	"github.com/ratheeshkumar25/blog-crud-api/config"
	"github.com/ratheeshkumar25/blog-crud-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *config.Config) *gorm.DB {
	if config.Database_url == "" {
		log.Fatal("Database URL is not set in the configuration")
	}
	// Log the connection details
	log.Printf("Connecting to DB with URL: %s", config.Database_url)

	DB, err := gorm.Open(postgres.Open(config.Database_url), &gorm.Config{})
	if err != nil {
		log.Fatal("connection to the database failed:", err)
	}

	err = DB.AutoMigrate(
		&model.BlogPost{},
	)

	if err != nil {
		log.Printf("error while migrating %v", err.Error())
		return nil
	}
	return DB
}
