package config

import (
	"log"
	"os" // 1. Impor package 'os' untuk mengakses environment variables
	"github.com/joho/godotenv" // 2. Impor package 'godotenv'
)
func Get() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error when load file configuration: ", err.Error())
	}
	return &Config{
		Server : Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database : Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Tz: os.Getenv("DB_TZ"),
		},
	}
}
