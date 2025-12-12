package config

import (
	// "log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort     string
	PostgresDsn string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		AppPort:     getEnv("APP_PORT", ":8080"),
		PostgresDsn: getEnv("POSTGRES_DSN", "host=localhost user=postgres password=postgres dbname=demo port=5432 sslmode=disable"),
	}
}

func getEnv(k, d string) string {
	if val, ok := os.LookupEnv(k); ok {
		return val
	}
	return d
}
