package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetMongoURI() string {
	catchError := godotenv.Load()
	if catchError != nil {
		log.Fatal("Couldn't load the .env file")
	}

	return os.Getenv("MONGOURI")
}
