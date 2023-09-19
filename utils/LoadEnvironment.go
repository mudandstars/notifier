package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	error := godotenv.Load()

	if error != nil {
		log.Fatalf("Error loading .env file")
	}

}
