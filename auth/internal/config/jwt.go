package config

import (

	"os"
	"time"

	"github.com/joho/godotenv"
)

type JWTConfig struct {
	AccessSecret  string
	RefreshSecret string
	AccessTTL     time.Duration
	RefreshTTL    time.Duration
}

var JWT *JWTConfig

func InitJWT() {
	_ = godotenv.Load()

	accessTTL, _ := time.ParseDuration(getEnv("JWT_ACCESS_TTL", "15m"))
	refreshTTL, _ := time.ParseDuration(getEnv("JWT_REFRESH_TTL", "720h"))

	JWT = &JWTConfig{
		AccessSecret:  getEnv("JWT_ACCESS_SECRET", "default_access_secret"),
		RefreshSecret: getEnv("JWT_REFRESH_SECRET", "default_refresh_secret"),
		AccessTTL:     accessTTL,
		RefreshTTL:    refreshTTL,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
