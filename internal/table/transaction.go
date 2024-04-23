package table

// Transaction -.
type Transaction struct {
	ID          int `gorm:"column:id"`
	SenderID    int `gorm:"column:sender_id"`
	RecipientID int `gorm:"column:recipient_id"`
	Amount      int `gorm:"column:amount"`
}
