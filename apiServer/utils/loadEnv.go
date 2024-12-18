package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	DatabaseUrl       string
	PortNumber        string
	AccessTokenSecret string
	Secret            string
}

func LoadEnvFiles() EnvVariables {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("Database url not found in the environment variable")
	}

	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		log.Fatal("Port number not found in the environment variable")
	}

	accessTokenSecret := os.Getenv("ACCESSTOKENSECRET")
	if accessTokenSecret == "" {
		log.Fatal("Access token secret not found in the environment variable")
	}
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("Secret not found in the environment variable")
	}

	return EnvVariables{
		DatabaseUrl:       databaseUrl,
		PortNumber:        portNumber,
		AccessTokenSecret: accessTokenSecret,
		Secret:            secret,
	}
}
