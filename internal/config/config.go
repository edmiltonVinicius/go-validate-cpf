package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var Global *config

type config struct {
	ENV            string
	PORT           string
	DB             *gorm.DB
	DB_HOST        string
	DB_PORT        string
	DB_PASSWORD    string
	DB_USER        string
	DB_NAME        string
	SSL_MODE       string
	RUN_MIGRATIONS bool
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Global = &config{
		ENV:            os.Getenv("ENV"),
		PORT:           os.Getenv("PORT"),
		DB_HOST:        os.Getenv("DB_HOST"),
		DB_PORT:        os.Getenv("DB_PORT"),
		DB_PASSWORD:    os.Getenv("DB_PASSWORD"),
		DB_USER:        os.Getenv("DB_USER"),
		DB_NAME:        os.Getenv("DB_NAME"),
		SSL_MODE:       os.Getenv("SSL_MODE"),
		RUN_MIGRATIONS: os.Getenv("RUN_MIGRATIONS") == "true",
	}
}
