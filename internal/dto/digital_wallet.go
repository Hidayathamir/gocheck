package dto

// ReqDigitalWalletTransfer represents the request data structure for transfer.
type ReqDigitalWalletTransfer struct {
	SenderID    uint `validate:"required,nefield=RecipientID"`
	RecipientID uint `validate:"required"`
	Amount      int  `validate:"required"`
}

// ResDigitalWalletTransfer represents the response data structure for transfer.
type ResDigitalWalletTransfer struct {
	ID uint
}
