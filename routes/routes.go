package routes

import (
	"github.com/Oluwaseun241/wallet/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
  app.Get("/", controllers.GetUser)
  app.Post("/api/auth/register", controllers.NewUser)
  //auth routes
  app.Post("/api/auth/login", controllers.LoginUser)
}
