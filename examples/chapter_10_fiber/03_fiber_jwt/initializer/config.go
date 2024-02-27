package initializer

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Secret = "secret"

// LoadEnvVariables функція для завантаження змінних середовища
func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Помилка при завантаженні .env файлу")
	}
	Secret = os.Getenv("SECRET")
}
