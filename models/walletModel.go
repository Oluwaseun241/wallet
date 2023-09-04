package models

import (
	"time"
  "gorm.io/gorm"
	"github.com/google/uuid"
  "math/rand"
)

type Wallet struct {
  gorm.Model
  ID  uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
  UserID  uuid.UUID `gorm:"type:uuid;not null"`
  WalletNumber  string  `gorm:"type:string;not null"`
  Balance float64 `gorm:"not null;default:0"`
  CreatedAt time.Time `gorm:"not null;autoCreateTime"`
  UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}

type WalletResponse struct {
	ID  uuid.UUID `json:"id,omitempty"`
	UserID  uuid.UUID `json:"user_id,omitempty"`
	WalletNumber  string  `json:"wallet_number,omitempty"`
  Balance int `json:"balance,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FundWallet struct {
  ID  uuid.UUID
  UserID  uuid.UUID
  Funds float64
  CreatedAt time.Time
}

func (Wallet) TableName() string {
  return "wallets"
}

func GenerateWallet() string {
  const charset = "0123456789"
  const walletNumberLength = 10
  seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	walletNumber := make([]byte, walletNumberLength)
	for i := range walletNumber {
		walletNumber[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(walletNumber)
}
