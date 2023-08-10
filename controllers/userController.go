package controllers

import (
  db "github.com/Oluwaseun241/wallet/config"
	"github.com/Oluwaseun241/wallet/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//Get all users
func GetUser(c *fiber.Ctx) error {
  var users []Models.User
  db.DB.Find(&users)
  return c.Status(fiber.StatusOK).JSON(users)
}

// Create New User
func NewUser(c *fiber.Ctx) error {
  user := &Models.User{
    ID: uuid.New(),
  }
  if err := c.BodyParser(user); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Invalid request format",
    })
  }

  // Validate data(db)
  user.Validate(db.DB)
  if db.DB.Error != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Validation failed",
      "details": db.DB.Error.Error(),
    })
  }

  result := db.DB.Create(&user)
  if result.Error != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to create user",
    })
  }
  return c.Status(fiber.StatusCreated).JSON(user)
}	
