package entity

import "time"

// Transaction represents entity transaction.
type Transaction struct {
	ID uint `gorm:"primarykey" json:"id"`

	SenderID    uint `gorm:"column:sender_id"    json:"sender_id"`
	RecipientID uint `gorm:"column:recipient_id" json:"recipient_id"`
	Amount      int  `gorm:"column:amount"       json:"amount"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
