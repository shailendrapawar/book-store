package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// 1: config for applicatiopn
type AppConfig struct {
	Env           string // dev or prod
	Port          string //eg:8080
	MigrationPath string // eg:MIGRATION_PATH
}

// 2: dbconfig
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// 3: jwt config
type JWTConfig struct {
	Secret      string
	ExpiryHours int
}

// 4: cookies config

type CookieConfig struct {
	HttpOnly    bool
	Secure      bool
	Domain      string
	ExpiryHours int
}

type Config struct {
	App    AppConfig
	DB     DBConfig
	JWT    JWTConfig
	Cookie CookieConfig
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file , using default environment/")
	}

	return &Config{
		App: AppConfig{
			Env:           getEnv("APP_ENV", "development"),
			Port:          getEnv("APP_PORT", "8080"),
			MigrationPath: getEnv("MIGRATIONS_PATH", "internal/db/migrations"),
		},
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "uk04ac2006"),
			Name:     getEnv("DB_NAME", "book_store"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:      getEnv("JWT_SECRET", "secret"),
			ExpiryHours: getEnvInt("JWT_EXPIRY_HOURS", 72),
		},
		Cookie: CookieConfig{
			HttpOnly:    true,
			Secure:      getEnv("APP_ENV", "development") == "production",
			Domain:      getEnv("COOKIE_DOMAIN", ""),
			ExpiryHours: getEnvInt("COOKIE_EXPIRY", 72),
		},
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v

	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}
