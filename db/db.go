package db

import (
	"fmt"
	"os"

	"github.com/yagyansh5/simplebank/models"
	"github.com/yagyansh5/simplebank/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), // postgres
		"5432",
		os.Getenv("DB_NAME"),
	)

	utils.Logger.Info(dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Logger.Fatalf("❌ Failed to connect to database: %v", err)
	}

	utils.Logger.Info("✅ Connected to PostgreSQL successfully!")

	DB.AutoMigrate(&models.Account{})
	DB.AutoMigrate(&models.Entry{})
	DB.AutoMigrate(&models.Transfer{})
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
