package presenters

import (
	"experiment/adapters/presenters/input"
	"experiment/adapters/presenters/output"
)

type CreateWalletPresenter interface {
	Present(message string) *CreateWalletResponse
}

type createWalletPresenter struct{}

func NewCreateWalletPresenter() CreateWalletPresenter {
	return &createWalletPresenter{}
}
func (cwp *createWalletPresenter) Present(wallet string) *CreateWalletResponse {
	return &CreateWalletResponse{Message: output.CreateWalletOutput{Message: wallet}}
}

type CreateWalletResponse struct {
	Message output.CreateWalletOutput `json:"message"`
}

type CreateWalletRequest struct {
	Wallet input.WalletInput `json:"wallet"`
}
