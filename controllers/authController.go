package controllers

import (
  "github.com/gofiber/fiber/v2"
  Models "github.com/Oluwaseun241/wallet/models"
  db "github.com/Oluwaseun241/wallet/config"
)

//User Login
func LoginUser(c *fiber.Ctx) error {
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
  if err := db.DB.Where("Email=?", userEmail).First(&user).Error; err != nil {
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
