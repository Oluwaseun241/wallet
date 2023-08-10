package controllers

import (
	"github.com/Oluwaseun241/wallet/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUser(c *fiber.Ctx, db *gorm.DB) error {
  var users []Models.User
  db.Find(&users)
  return c.Status(fiber.StatusOK).JSON(users)
}

// Create New User
func NewUser(c *fiber.Ctx, db *gorm.DB) error {
  //user := new(Models.User)
  user := &Models.User{
    ID: uuid.New(),
  }
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

func LoginUser(c *fiber.Ctx, db *gorm.DB) error {
  var loginReq Models.SignInInput
  if err := c.BodyParser(&loginReq); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "success": false,
      "message": "Invalid request format",
    })
  }
  userEmail := loginReq.Email
  userPassword := loginReq.Password

  var user Models.User
  if err := db.Where("Email=?", userEmail).First(&user).Error; err != nil {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
      "success": false,
      "message": "Invalid Credential",
    })
  }

  if user.Password != userPassword {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
      "success": false,
      "message": "Invalid Credential",
    })
  }

  return c.Status(fiber.StatusCreated).JSON(fiber.Map{
    "success": true,
    "message": "Success",
  })
}	
