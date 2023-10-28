package controllers

import (
	db "github.com/Oluwaseun241/wallet/config"
	Models "github.com/Oluwaseun241/wallet/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
  "gorm.io/gorm"
)

func TransferFund(c *fiber.Ctx) error {
  senderWalletNumber := c.Params("wallet_number")
  
  var senderWallet Models.Wallet
  if err := db.DB.Find(&senderWallet, "wallet_number=?", senderWalletNumber).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Sender's wallet not found",
			})
	}
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database error",
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
  if err := db.DB.Find(&receiverWallet, "wallet_number = ?", transaction.ReceiverID).First(&receiverWallet).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Receiver's wallet number not found",
			})
		}
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database error",
		})
	}

  if transaction.Amount <= 0 {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
        "error":   "Invalid transaction",
        "message": "Amount must be a positive number",
    })
  }

  if senderWallet.Balance < transaction.Amount {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Insufficient funds",
			"message": "Sender does not have enough balance for the transfer",
		})
  }

  transactionReference := Models.GenerateTransactionReference(senderWallet.UserID.String(), transaction.ReceiverID)

  transaction = Models.Transaction{
    ID: uuid.New(),
    SenderID: senderWalletNumber,
    ReceiverID: transaction.ReceiverID,
    Amount: transaction.Amount,
    TransactionReference: transactionReference,
  }
  
  // Save the transaction and wallet updates to the database
  tx := db.DB.Begin()
  if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create transaction",
		})
	}

  senderWallet.Balance -= transaction.Amount
  if err := tx.Save(&senderWallet).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update sender's wallet",
		})
	}
	
  receiverWallet.Balance += transaction.Amount
  if err := tx.Save(&receiverWallet).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update receiver's wallet",
		})
	}

  if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database transaction error",
		})
	}  

  return c.Status(fiber.StatusCreated).JSON(fiber.Map{
    "status": true,
    "message": "Transfer was sucessful", 
    "balance": senderWallet.Balance,
    "transaction_ref": transactionReference,
  })
}
