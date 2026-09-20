package presenters

import "experiment/adapters/presenters/output"

type GetWalletsByOwnerPresenter interface {
	Present(wallets []output.WalletOutput) *GetWalletsByOwnerResponse
}

type getWalletsByOwnerPresenter struct{}

func NewGetWalletsByOwnerPresenter() GetWalletsByOwnerPresenter {
	return &getWalletsByOwnerPresenter{}
}

func (gop *getWalletsByOwnerPresenter) Present(wallets []output.WalletOutput) *GetWalletsByOwnerResponse {
	return &GetWalletsByOwnerResponse{Wallets: wallets}
}

type GetWalletsByOwnerResponse struct {
	Wallets []output.WalletOutput `json:"wallets"`
}