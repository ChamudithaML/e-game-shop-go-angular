package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Declare variables to store the environment variables
var (
	API_GATEWAY_PORT  string
	FILE_SERVICE_URL  string
	GAME_SERVICE_URL  string
	SALES_SERVICE_URL string
	USER_SERVICE_URL  string
)

// LoadConfig loads the environment variables from the .env file
func LoadConfig() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set environment variables
	API_GATEWAY_PORT = os.Getenv("API_GATEWAY_PORT")
	FILE_SERVICE_URL = os.Getenv("FILE_SERVICE_URL")
	GAME_SERVICE_URL = os.Getenv("GAME_SERVICE_URL")
	SALES_SERVICE_URL = os.Getenv("SALES_SERVICE_URL")
	USER_SERVICE_URL = os.Getenv("USER_SERVICE_URL")

	// Optional: Print out to verify environment variables are loaded
	log.Println("Loaded environment variables")
}
