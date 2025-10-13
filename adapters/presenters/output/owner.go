package output

type CreateOwnerOutput struct {
	Message string `json:"message"`
}

type GetOwnerOutput struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Document string         `json:"document"`
	Wallets  []WalletOutput `json:"wallets"`
}

type GetOwnersOutput struct {
	Owners []GetOwnerOutput `json:"owners"`
}
