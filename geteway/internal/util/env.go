package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


func LoadEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env yuklanmadi, tizim env ishlatilmoqda...")
	}

	value := os.Getenv(key)
	if value == "" {
		log.Printf(" %s env topilmadi", key)
	}
	return value
}
