package gocheckhttp

// ReqDigitalWalletTransfer -.
type ReqDigitalWalletTransfer struct {
	SenderID    uint `json:"sender_id"`
	RecipientID uint `json:"recipient_id"`
	Amount      int  `json:"amount"`
}

// ResDigitalWalletTransfer -.
type ResDigitalWalletTransfer struct {
	Data  ResDataDigitalWalletTransfer `json:"data"`
	Error string                       `json:"error"`
}

// ResDataDigitalWalletTransfer -.
type ResDataDigitalWalletTransfer struct {
	ID uint `json:"id"`
}
