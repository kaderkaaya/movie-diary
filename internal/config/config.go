package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort    string
	AppEnv     string
	DbDsn      string
	JwtSecret  string
	TmdbApiKey string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		AppEnv:     getEnv("APP_ENV", "development"),
		DbDsn:      mustEnv("DB_DSN"),
		JwtSecret:  mustEnv("JWT_SECRET"),
		TmdbApiKey: mustEnv("TMDB_API_KEY"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing required env: %s", key)
	}
	return v
}
