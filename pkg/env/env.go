package env

import (
	"log"

	"github.com/joho/godotenv"
)

// Загрузка файла .env
func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}
