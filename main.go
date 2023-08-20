package main

import (
	"github.com/Oluwaseun241/wallet/config"
	"github.com/Oluwaseun241/wallet/routes"
	"github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
)


func main() {
  // db connection
  config.InitDB() 
  app := fiber.New()
  app.Use(cors.New())
  routes.Setup(app)

  app.Listen(":3000")
}
