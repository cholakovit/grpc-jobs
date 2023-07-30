package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(envVariableRequest string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envVariable := os.Getenv(envVariableRequest)
	if envVariable == "" {
		log.Fatal("The ENV variable is not set")
	}

	return envVariable
}