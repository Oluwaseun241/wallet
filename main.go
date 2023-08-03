package main

import (
	"fmt"

	"github.com/Oluwaseun241/wallet.git/models"
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

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello")
  })

  app.Listen(":3000")
}
