package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Global *config

type config struct {
	ENV  string
	PORT string
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Global = &config{
		ENV:  os.Getenv("ENV"),
		PORT: os.Getenv("PORT"),
	}
}
