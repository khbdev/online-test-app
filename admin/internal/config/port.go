package config

import (

	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	GRPCPort string
}

var Server *ServerConfig

func InitPort() {
	_ = godotenv.Load()

	Server = &ServerConfig{
		GRPCPort: getPortenv("GRPC_PORT", ":50052"),
	}
}

func getPortenv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}