
package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)


func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
