package main

import (

  "github.com/Oluwaseun241/wallet/routes"	
	"github.com/gofiber/fiber/v2"
  "github.com/Oluwaseun241/wallet/config"
)


func main() {
  // db connection
  config.InitDB() 
  app := fiber.New()
  routes.Setup(app)

  app.Listen(":3000")
}
