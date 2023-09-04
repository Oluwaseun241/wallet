package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
  gorm.Model
  ID  uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
  SenderID  uuid.UUID `gorm:"type:uuid;not null"`
  ReceiverID  uuid.UUID `gorm:"type:uuid;not null"`
  Amount  float64 `gorm:"not null"`
  CreatedAt time.Time `gorm:"not null;autoCreateTime"`
  UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}


func (Transaction) TableName() string {
  return "transactions"
}
