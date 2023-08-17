package controllers

import (

	"github.com/Oluwaseun241/wallet/config"
	db "github.com/Oluwaseun241/wallet/config"
	Models "github.com/Oluwaseun241/wallet/models"
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
  // Auth
  user, err := Models.Authenticate(db.DB, loginReq.Email, loginReq.Password)
  if err != nil {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid credentials",
		})
  }
  
  // JWT
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

// Logout
func LogoutUser() {

}
