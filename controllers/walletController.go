package controllers

import (
	"github.com/Oluwaseun241/wallet/models"
  db "github.com/Oluwaseun241/wallet/config"
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
  wallet_number := models.GenerateWallet() 
  wallet := &Models.Wallet{
    ID: uuid.New(),
    UserID: userID,
    WalletNumber: wallet_number,
  }
  if err := c.BodyParser(wallet); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Invalid request format",
    })
  }

  var existingWallet Models.Wallet
  if err := db.DB.Where("user_id = ?", userID).First(&existingWallet).Error; err == nil {
    return c.Status(fiber.StatusConflict).JSON(fiber.Map{
      "error": "User already has a wallet",
    })
  }

  result := db.DB.Create(&wallet)
  if result.Error != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to create wallet",
    })
  }

  // Response
  walletResponse := Models.WalletResponse{
    ID: wallet.ID,
    UserID: wallet.UserID,
    WalletNumber: wallet.WalletNumber,
    Balance: int(wallet.Balance),
    CreatedAt: wallet.CreatedAt,
    UpdatedAt: wallet.UpdatedAt,
}
  return c.Status(fiber.StatusCreated).JSON(walletResponse)
}
