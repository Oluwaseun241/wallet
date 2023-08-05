package controllers

import (
  "gorm.io/gorm"
  "github.com/Oluwaseun241/wallet/Models"
  "github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx, db *gorm.DB) error {
  var users []Models.User
  db.Find(&users)
  return c.Status(200).JSON(users)
}

func NewUser(c *fiber.Ctx, db *gorm.DB) error {
  user := new(Models.User)
  if err := c.BodyParser(user); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Invalid request format",
    })
  }

  // Validate
  user.Validate(db)
  if db.Error != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Validation failed",
      "details": db.Error.Error(),
    })
  }

  result := db.Create(&user)
  if result.Error != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to create user",
    })
  }
  return c.Status(fiber.StatusCreated).JSON(user)
}
