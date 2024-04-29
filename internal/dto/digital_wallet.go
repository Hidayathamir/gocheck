package dto

// ReqTransfer represents the request data structure for transfer.
type ReqTransfer struct {
	SenderID    uint
	RecipientID uint
	Amount      int
}

// ResTransfer represents the response data structure for transfer.
type ResTransfer struct {
	ID uint
}
