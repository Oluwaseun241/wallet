package main

import (
	"fmt"

	"github.com/Oluwaseun241/wallet.git/Models"
	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func initDB(){
  db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect to DB")
  }
  fmt.Println("DB connected")
  db.AutoMigrate(&models.User{})
  fmt.Println("DB migrated")
}

func main() {
  app := fiber.New()
  
  initDB()
  app.Use(app)

  app.Listen(":3000")
}
