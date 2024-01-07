package env

import (
	"fmt"

	"github.com/joho/godotenv"
)

// Загрузка файла .env
func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
