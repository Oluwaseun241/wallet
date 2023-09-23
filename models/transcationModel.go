package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
  gorm.Model
  ID  uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
  SenderID  string `gorm:"type:string;not null"`
  ReceiverID  string `gorm:"type:string;not null"`
  Amount  float64 `gorm:"not null"`
  CreatedAt time.Time `gorm:"not null;autoCreateTime"`
  UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}


type TranscationResponse struct {
  ID  uuid.UUID `json:"id,omitempty"`
	SenderID  string `json:"sender_id,omitempty"`
	ReceiverID  string `json:"receiver_id,omitempty"`
  Amount float64 `json:"amount,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Transaction) TableName() string {
  return "transactions"
}
