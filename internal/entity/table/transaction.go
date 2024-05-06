package table

// Transaction represents schema table transactions.
var Transaction = transaction{
	ID:          "id",
	SenderID:    "sender_id",
	RecipientID: "recipient_id",
	Amount:      "amount",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	Dot: &transaction{
		ID:          "transactions.id",
		SenderID:    "transactions.sender_id",
		RecipientID: "transactions.recipient_id",
		Amount:      "transactions.amount",
		CreatedAt:   "transactions.created_at",
		UpdatedAt:   "transactions.updated_at",
	},
}

type transaction struct {
	ID string

	SenderID    string
	RecipientID string
	Amount      string

	CreatedAt string
	UpdatedAt string

	Dot *transaction
}

func (transaction) TableName() string {
	return "transactions"
}
