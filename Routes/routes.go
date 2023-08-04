package Routes

import (
	"github.com/gofiber/fiber/v2"
  "github.com/Oluwaseun241/wallet/Controllers"
  "gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
  app.Get("/", func(c *fiber.Ctx) error {
    return Controllers.GetUser(c,db)
  })
  app.Post("/", func(c *fiber.Ctx) error {
    return Controllers.NewUser(c,db)
  })
}
