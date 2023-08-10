package controllers

import (
  "gorm.io/gorm"
  "github.com/gofiber/fiber/v2"
  "github.com/Oluwaseun241/wallet/models"
)

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
