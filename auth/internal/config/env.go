package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type GRPCConfig struct {
	Host string
	Port string
}

var GRPC *GRPCConfig

func InitGRPC() {
	_ = godotenv.Load()

	host := getEnvPORT("GRPC_HOST", "127.0.0.1")
	port := getEnvPORT("GRPC_PORT_ADMIN", "50052")

	GRPC = &GRPCConfig{
		Host: host,
		Port: port,
	}
}


func (g *GRPCConfig) Address() string {
	return fmt.Sprintf("%s:%s", g.Host, g.Port)
}


func getEnvPORT(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}