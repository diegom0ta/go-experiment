package output

type CreateWalletOutput struct {
	Message string `json:"message"`
}

type WalletOutput struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
