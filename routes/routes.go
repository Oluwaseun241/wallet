package routes

import (
	"github.com/Oluwaseun241/wallet/controllers"
	"github.com/gofiber/fiber/v2"
  //"gorm.io/gorm"
)

func Setup(app *fiber.App) {
  // users routes
  app.Get("/api/users", controllers.GetUser)
  app.Get("/api/users/:userId", controllers.GetUserId)
  app.Put("/api/users/:userId", controllers.UpdateUser)
  app.Post("/api/auth/register", controllers.NewUser)
  //auth routes
  app.Post("/api/auth/login", controllers.LoginUser)
}
