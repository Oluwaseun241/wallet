package controllers

import (
	"github.com/Oluwaseun241/wallet/models"
  db "github.com/Oluwaseun241/wallet/config"
	Models "github.com/Oluwaseun241/wallet/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetWallet(c *fiber.Ctx) error {
  userIDStr := c.Params("userId")
  
  // Convert the userId string to a UUID
  userID, err := uuid.Parse(userIDStr)
  if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Invalid userId format",
    })
  }
  
  // Find the wallet associated with the user by userID
  var wallet Models.Wallet
  if err := db.DB.First(&wallet, "user_id = ?", userID).Error; err != nil {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "Wallet not found",
    })
  }
  return c.Status(fiber.StatusOK).JSON(wallet)
}

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

func FundWallet(c *fiber.Ctx) error {
  wallet_number := c.Params("wallet_number")
  var wallet Models.Wallet

  result := db.DB.Find(&wallet, "wallet_number=?", wallet_number)
  if result.Error != nil {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "Wallet number not found",
    })
  }
  
  var updatedWallet Models.Wallet
  if err := c.BodyParser(&updatedWallet); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Invalid request format",
    })
  }

  wallet.Balance += updatedWallet.Balance

  result = db.DB.Save(&wallet)
  if result.Error != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to update wallet",
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "Wallet funding sucessfully",
    "success": true,
    "balance": wallet.Balance,
  })
}

func WihdrawFund(c *fiber.Ctx) error {
  wallet_number := c.Params("wallet_number")
  var wallet Models.Wallet

  result := db.DB.Find(&wallet, "wallet_number=?", wallet_number)
  if result.Error != nil {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "Wallet number not found",
    })
  }
  
  var updatedWallet Models.Wallet
  if err := c.BodyParser(&updatedWallet); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Invalid request format",
    })
  }

  if wallet.Balance < updatedWallet.Balance {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Insufficient funds",
    })
  }

  wallet.Balance -= updatedWallet.Balance 

  result = db.DB.Save(&wallet)
  if result.Error != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to update wallet",
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "Withdrawal sucessfully",
    "success": true,
    "balance": wallet.Balance,
  }) 
}
