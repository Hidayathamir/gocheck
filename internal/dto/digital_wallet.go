package dto

// ReqTransfer -.
type ReqTransfer struct {
	SenderID    uint
	RecipientID uint
	Amount      int
}

// ResTransfer -.
type ResTransfer struct {
	ID uint
}
