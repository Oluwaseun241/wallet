package controllers

import (
	db "github.com/Oluwaseun241/wallet/config"
	Models "github.com/Oluwaseun241/wallet/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func TransferFund(c *fiber.Ctx) error {
  senderWalletNumber := c.Params("wallet_number")
  
  var senderWallet Models.Wallet
  if err := db.DB.Find(&senderWallet, "wallet_number=?", senderWalletNumber).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Sender's wallet not found",
		})
	}

  var transaction Models.Transaction 
  if err := c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

  if senderWalletNumber == transaction.ReceiverID {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid transaction",
			"message": "Cannot transfer to yourself",
		})
  }

  // Check if the receiver's wallet exists
  var receiverWallet Models.Wallet
  if err := db.DB.Find(&receiverWallet, "wallet_number=?", transaction.ReceiverID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Receiver's wallet number not found",
		})
	}

  if senderWallet.Balance < transaction.Amount {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Insufficient funds",
			"message": "Sender does not have enough balance for the transfer",
		})
  }

  transaction = Models.Transaction{
    ID: uuid.New(),
    SenderID: senderWalletNumber,
    ReceiverID: transaction.ReceiverID,
    Amount: transaction.Amount,
  }
  
  senderWallet.Balance -= transaction.Amount
	receiverWallet.Balance += transaction.Amount

  return c.Status(fiber.StatusCreated).JSON(transaction)
}
