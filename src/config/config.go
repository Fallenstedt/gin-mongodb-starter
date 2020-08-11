package config 

import (
	"log"
	"sync"
	"os"
	"github.com/joho/godotenv"
)

type Config struct { 
	DbName string
	Mode 			string
	MongoConnection string

}

var once sync.Once
var instance *Config

func Get() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		
		config := Config {
			Mode: os.Getenv("MODE"),
			MongoConnection: os.Getenv("MONGO_CONNECTION"),
			DbName: os.Getenv("DB_NAME"),
		}
		instance = &config
	})
	return instance
}