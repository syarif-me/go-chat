package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	env := os.Getenv("APP_ENV")

	if env == "" {
		err := godotenv.Load()

		if err != nil {
			log.Println(".env file not found")
		}
	}
}
