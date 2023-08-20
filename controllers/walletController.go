package controllers

import (
	Models "github.com/Oluwaseun241/wallet/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewWallet(c *fiber.Ctx) error {
  userIDStr := c.Locals("userID").(string)
  
  // Convert to uuid
  userID, err := uuid.Parse(userIDStr)
  if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Invalid user ID format",
    })
  }

  wallet := &Models.Wallet{
    ID: uuid.New(),
    UserID: userID,
  }
  
  if err := c.BodyParser(wallet); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Invalid request format",
    })
  }
  return c.Status(fiber.StatusCreated).JSON(wallet)
}
