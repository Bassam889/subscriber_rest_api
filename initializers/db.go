package initializers

import (
	"fmt"
	"os"

	"app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faield to connect to database")
	} else {
		fmt.Println("Successfully connected to database")
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Subscriber{})
}
