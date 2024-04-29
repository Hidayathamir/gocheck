package dto

// ReqDigitalWalletTransfer represents the request data structure for transfer.
type ReqDigitalWalletTransfer struct {
	SenderID    uint
	RecipientID uint
	Amount      int
}

// ResDigitalWalletTransfer represents the response data structure for transfer.
type ResDigitalWalletTransfer struct {
	ID uint
}
