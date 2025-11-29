package config

import (
	"fmt"
	"log"
	"os"
	"test-section-service/internal/models"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB 

func InitDB() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading environment variables directly")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)


	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get db instance: %v", err)
	}

	
	sqlDB.SetMaxIdleConns(10)           
	sqlDB.SetMaxOpenConns(100)          
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf(" Database ping failed: %v", err)
	}
DB = db


	if err := DB.AutoMigrate(&models.Section{}, models.Question{}, models.Option{}); err != nil {
		log.Fatalf(" AutoMigrate error: %v", err)
	}

	log.Println("✅ Database connected and migrated successfully!")
}