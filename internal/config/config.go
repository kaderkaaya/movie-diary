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
	OmdbApiKey string
}

func Load() *Config {
	_ = godotenv.Load()
	//dosyasındaki key-value’ları alıp process environment’a yüklemek
	//burda büyük structları kopyalamak yerine pointer ile return ediyoruz.
	//bu sayede bellek kullanımı azalır ve performans artar.
	//dependency injection yapıyoruz ve böylece daha clean bir yapı olur.

	return &Config{
		// burda &Config yaparak Tüm config değerlerini tek bir struct içinde topluyoruz.
		//&Config → config’in adresini gönderiyorsun
		//* ile de o adresteki değeri değiştiriyorsun.
		//burda kopya değil direkt configi alıyoruz.
		AppPort:    getEnv("APP_PORT", "8080"),
		AppEnv:     getEnv("APP_ENV", "development"),
		DbDsn:      mustEnv("DB_DSN"),
		JwtSecret:  mustEnv("JWT_SECRET"),
		TmdbApiKey: mustEnv("TMDB_API_KEY"),
		OmdbApiKey: mustEnv("OMDB_API_KEY"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// Burda mustenv fonksiyonu ile env dosyasındaki key’leri alıyoruz ve eğer boşsa uygulamayı baslatmadan öldürür.
func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing required env: %s", key)
	}
	return v
}
