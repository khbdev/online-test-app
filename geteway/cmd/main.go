package main

import (
	"geteway-service/internal/handler"
	"geteway-service/internal/middleware"
	"geteway-service/internal/util"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	handler.SetupRoutes(r)


	port := util.LoadEnv("PORT")
	if port == "" {
		port = ":8081"
	}
	if port[0] != ':' {
		port = ":" + port
	}

	log.Printf(" Gateway server %s portda ishlayapti...", port)
	if err := r.Run(port); err != nil {
		log.Fatalf(" Server ishga tushmadi: %v", err)
	}
}
