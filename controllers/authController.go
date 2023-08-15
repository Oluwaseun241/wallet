package controllers

import (

	db "github.com/Oluwaseun241/wallet/config"
	Models "github.com/Oluwaseun241/wallet/models"
  "github.com/Oluwaseun241/wallet/config"
	"github.com/gofiber/fiber/v2"
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
  userID := user.ID
  token, err := config.GenerateToken(userID)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "success": false,
      "message": "Failed to generate token",
    })
  }
  return c.Status(fiber.StatusCreated).JSON(fiber.Map{
    "success": true,
    "message": "Success",
    "token": token,
  })
}
