package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser    string
	DBPass    string
	DBHost    string
	DBPort    string
	DBName    string
	GinMode   string
	Port      string
	JwtSecret string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("⚠️ No .env file found, using system envs")
	}

	AppConfig = &Config{
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASS"),
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_NAME"),
		GinMode:   os.Getenv("GIN_MODE"),
		Port:      os.Getenv("PORT"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}
