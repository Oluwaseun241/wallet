package controllers

import (
	"github.com/Oluwaseun241/wallet/auth"
	db "github.com/Oluwaseun241/wallet/config"
	Models "github.com/Oluwaseun241/wallet/models"
	"github.com/gofiber/fiber/v2"
  "github.com/google/uuid"
)

// User Login
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
  username := user.Name
  token, err := auth.GenerateToken(userID, username)
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

// Token refresh
func ResfreshToken(c *fiber.Ctx) error {
  oldToken := c.Get("Authorization")

  claims, err := auth.ValidateToken(oldToken)
  if err != nil {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
  }
  
  userIDStr, ok := claims["sub"].(string)
  if !ok {
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
          "error": "Invalid user ID format",
      })
  }

  userID, err := uuid.Parse(userIDStr)
  if err != nil {
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
          "error": "Failed to parse user ID",
      })
  }

  user := Models.User{}
  username := user.Name

  newToken, err := auth.GenerateToken(userID, username)
  if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate new token",
		})
	}

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": newToken,
	})
}

// Logout
func LogoutUser() {
  
}
