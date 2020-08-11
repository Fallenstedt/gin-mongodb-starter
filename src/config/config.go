package config 

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct { 
	Mode 			string
	MongoConnection string
}

func Init() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config {
		Mode: os.Getenv("MODE"),
		MongoConnection: os.Getenv("MONGO_CONNECTION"),
	}

	return &config, nil
}