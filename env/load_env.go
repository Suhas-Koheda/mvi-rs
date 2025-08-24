package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var EnvKey string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
		log.Println("Using system environment variables instead")
	}

	EnvKey = os.Getenv("ENVKEY")
	if EnvKey == "" {
		log.Println("Warning: ENVKEY environment variable is not set")
	}
}
