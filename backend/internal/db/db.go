package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Subhadip006/scalura/internal/config"
	"github.com/Subhadip006/scalura/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	config.LoadEnv()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to database")

	err = db.AutoMigrate(models.User{},
		models.Task{},
		models.Goal{},
		models.Skill{},
	)

	if err != nil {
		log.Println("Failed to migrate models")
	} else {
		log.Println("migration successful")
	}
	DB = db
}
