package main

import (
	"fmt"

	"github.com/Oluwaseun241/wallet/Models"
  routes "github.com/Oluwaseun241/wallet/Routes"
	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
  //"github.com/qor/validations"
)

func initDB() *gorm.DB{
  db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
  //validations.RegisterCallbacks(db)
  if err != nil {
    panic("failed to connect to DB")
  }
  fmt.Println("DB connected")
  db.AutoMigrate(&Models.User{})
  fmt.Println("DB migrated")
  return db
}

func main() {
  db := initDB()
  app := fiber.New()
  routes.Setup(app,db)

  app.Listen(":3000")
}
