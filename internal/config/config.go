package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DbUrl string
}

func Load() Config{
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	
	return Config{
		Port: getEnv("PORT", "8080"),
		DbUrl: getEnv("DB_URL", ""),
	}
}

func getEnv(key string, fallback string) string{

	if value,ok := os.LookupEnv(key); ok {
		return value;
	}

	return fallback;
}