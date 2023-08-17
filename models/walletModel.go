package models

import (
	"time"
  "gorm.io/gorm"
	"github.com/google/uuid"
)

type Wallet struct {
  gorm.Model
  ID  uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
  UserID  uuid.UUID `gorm:"type:uuid;not null"`
  Balance float64 `gorm:"not null;default:0"`
  CreatedAt time.Time `gorm:"not null;autoCreateTime"`
  UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}

func (Wallet) TableName() string {
  return "wallets"
}
