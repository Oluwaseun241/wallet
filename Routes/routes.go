package Routes

import "github.com/gofiber/fiber/v2"

func setup(app *fiber.App) {
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello")
  })
}
