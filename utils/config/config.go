package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		PostgresUser     string
		PostgresPassword string
		PostgresHost     string
		PostgresDb       string
	}
)

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error on running godotenv")
	}

	config := &Config{
		PostgresUser:     GetEnv("POSTGRES_USER", "admin"),
		PostgresPassword: GetEnv("POSTGRES_PASSWORD", "admin"),
		PostgresHost:     GetEnv("POSTGRES_HOST", "localhost"),
		PostgresDb:       GetEnv("POSTGRES_DB", "admin"),
	}

	return config
}

func GetEnv(key, val string) string {
	if os.Getenv(key) == "" {
		return val
	}

	return os.Getenv(key)
}
