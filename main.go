package main

import (

  routes "github.com/Oluwaseun241/wallet/Routes"	
	"github.com/gofiber/fiber/v2"
  "github.com/Oluwaseun241/wallet/config"
)


func main() {
  db := config.InitDB() 
  app := fiber.New()
  routes.Setup(app,db)

  app.Listen(":3000")
}
