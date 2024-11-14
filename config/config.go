package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
    // Memuat file .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

// Untuk mendapatkan variabel lingkungan dari .env
func GetEnv(key string, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        return defaultValue
    }
    return value
}
