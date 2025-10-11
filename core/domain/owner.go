package domain

type Owner struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Document string   `json:"document"`
	Wallets  []Wallet `json:"wallets"`
}
