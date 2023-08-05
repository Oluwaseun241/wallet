package Routes

import (
	"github.com/gofiber/fiber/v2"
  "github.com/Oluwaseun241/wallet/Controllers"
  "gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
  app.Get("/", func(c *fiber.Ctx) error {
    return controllers.GetUser(c,db)
  })
  app.Post("/api/auth/register", func(c *fiber.Ctx) error {
    return controllers.NewUser(c,db)
  })
  app.Post("/api/auth/login", func(c *fiber.Ctx) error {
    return controllers.LoginUser(c,db)
  })
}
