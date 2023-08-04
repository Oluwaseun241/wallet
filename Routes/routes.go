package Routes

import (
	"github.com/gofiber/fiber/v2"
  "github.com/Oluwaseun241/wallet/Models"
  "gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
  app.Get("/", func(c *fiber.Ctx) error {
    return Models.GetUser(c,db)
  })
  app.Post("/", func(c *fiber.Ctx) error {
    return Models.NewUser(c,db)
  })
}
