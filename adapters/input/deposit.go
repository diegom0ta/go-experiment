package input

type DepositInput struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
