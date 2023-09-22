package controllers

import (
	db "github.com/Oluwaseun241/wallet/config"
	Models "github.com/Oluwaseun241/wallet/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func TransferFund(c *fiber.Ctx) error { 
  wallet_number := c.Params("wallet_number")
  var wallet Models.Wallet

  senderWallet := db.DB.Find(&wallet, "wallet_number=?", wallet_number)
  if senderWallet.Error != nil {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "Wallet number not found",
    })
  }
  
  transaction := &Models.Transaction{
    ID: uuid.New(),
    SenderID: wallet_number,
  }
  
  receiverWallet := db.DB.Find(&wallet, "wallet_number=?", transaction.ReceiverID)
  if receiverWallet.Error != nil {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "Receiver's wallet number not found",
    })
  }

  if transaction.SenderID == transaction.ReceiverID {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "Invalid transaction",
      "message": "Cannot transfer to yourself",
    })
  }

  result := db.DB.Save(&transaction)
  if result.Error != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed transaction",
    })
  }

  return c.Status(fiber.StatusCreated).JSON(transaction)
}
