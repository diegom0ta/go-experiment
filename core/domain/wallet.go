package domain

type Wallet struct {
	ID         string `json:"id"`
	WalletName string `json:"name"`
	Balance    int    `json:"balance"`
	OwnerID    string `json:"owner_id"`
}
