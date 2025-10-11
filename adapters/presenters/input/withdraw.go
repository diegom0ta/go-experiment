package input

type WithdrawInput struct {
	Amount   float64     `json:"amount"`
	Currency string      `json:"currency"`
	Wallet   WalletInput `json:"wallet"`
}
