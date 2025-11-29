package config

import (
	"os"
)

func GetEnv(key string) *string {
	val, exists := os.LookupEnv(key)
	if !exists || val == "" {
		return nil
	}
	return &val
}
