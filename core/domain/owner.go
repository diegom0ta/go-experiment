package domain

type Owner struct {
	ID       string
	Name     string
	Email    string
	Document string
	Wallets  []Wallet
}
