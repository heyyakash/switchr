package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetString(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error Loading the .env file")
	}
	return os.Getenv(key)
}
