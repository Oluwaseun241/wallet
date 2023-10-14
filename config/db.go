package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
  "gorm.io/driver/postgres"

	Models "github.com/Oluwaseun241/wallet/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

  env := os.Getenv("APP_ENV")
  
  if env != "railway" {
    // Load env
    if err := godotenv.Load(); err != nil {
      log.Fatal("Error loading .env file")
    }
  } 
  
  dbHost := os.Getenv("DB_HOST")
  dbPort := os.Getenv("DB_PORT")
  dbUser := os.Getenv("DB_USER")
  dbPassword := os.Getenv("DB_PASSWORD")
  dbName := os.Getenv("DB_NAME")

  dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect to DB")
  }
  fmt.Println("DB connected")
  
  DB = db
  db.AutoMigrate(&Models.User{},&Models.Wallet{},&Models.Transaction{})
  fmt.Println("DB migrated")
}
