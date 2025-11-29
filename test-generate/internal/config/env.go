package config

import (
	"log"
	"os"
)

func GetEnv(key string) string {
    val, exists := os.LookupEnv(key)
    if !exists || val == "" {
        log.Fatalf("Environment variable %s not set", key)
    }
    return val
}