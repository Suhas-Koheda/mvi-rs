package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var EnvKey string

func LoadEnv()  {
	err:=godotenv.Load()
	if err!=nil{
		log.Fatal("Failed to load .env file")
	}

	EnvKey=os.Getenv("envkey")
}
