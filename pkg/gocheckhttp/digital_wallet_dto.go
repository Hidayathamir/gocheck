package gocheckhttp

// ReqDigitalWalletTransfer represents the request structure for transfer.
type ReqDigitalWalletTransfer struct {
	RecipientID uint `json:"recipient_id"`
	Amount      int  `json:"amount"`
}

// ResDigitalWalletTransfer represents the response structure for transfer.
type ResDigitalWalletTransfer struct {
	Data  ResDataDigitalWalletTransfer `json:"data"`
	Error string                       `json:"error"`
}

// ResDataDigitalWalletTransfer represents the response data structure for transfer.
type ResDataDigitalWalletTransfer struct {
	ID uint `json:"id"`
}
