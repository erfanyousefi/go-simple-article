package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPass     string
	DBPort     string
	DBName     string
	DBHost     string
	Driver     string
	ServerPort string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error: not found env file for load")
	}
	cfg := Config{
		DBUser:     getEnv("DB_USER", "root"),
		DBPass:     getEnv("DB_PASS", "root"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBName:     getEnv("DB_NAME", "db_article"),
		Driver:     getEnv("DB_DRIVER", "mysql"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
	return cfg
}

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
