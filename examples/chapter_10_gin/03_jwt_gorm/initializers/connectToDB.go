package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"jwt_gorm/models"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB")
	}
}

func SyncDatabase() error {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("failed to sync database: %w", err)
	}
	err = DB.AutoMigrate(&models.Profile{})
	if err != nil {
		return fmt.Errorf("failed to sync database: %w", err)
	}
	return nil
}
