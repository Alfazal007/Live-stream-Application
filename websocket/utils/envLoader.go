package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	Secret string
}

func LoadEnvFiles() EnvVariables {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("Secret not found in the environment variable")
	}

	return EnvVariables{
		Secret: secret,
	}
}
