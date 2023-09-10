package main

import (
	"log"
	"os"

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
  port := os.Getenv("PORT")
  if port == "" {
    port = ":3000"
  } else {
    port = ":" + port
  }
  err := app.Listen(port)
  if err != nil {
    log.Fatalf("Error starting server: %v", err)
  }
}
