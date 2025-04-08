package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if os.Getenv("BUILD_TYPE") == "dockerfile" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("❌ Error loading .env file. Using system environment variables.")
		} else {
			fmt.Println("✅ .env file loaded successfully")
		}
	} else {
		fmt.Println("🔸 Running in Docker. Using system environment variables.")
	}
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
