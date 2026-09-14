package controllers

import (
	"context"
	"experiment/adapters/presenters/input"
	"experiment/core/domain"
	"experiment/usecases"
	"uuid"
)

type CreateWalletController interface {
	HandleCreateWallet(ctx context.Context, email string, wallet *input.WalletInput) error
}

type createWalletController struct {
	createWalletUseCase usecases.ICreateWalletUseCase
}

func NewCreateWalletController(cwu usecases.ICreateWalletUseCase) CreateWalletController {
	return &createWalletController{createWalletUseCase: cwu}
}

func (coc *createWalletController) HandleCreateWallet(ctx context.Context, email string, wallet *input.WalletInput) error {
	id := uuid.NewV7().String()

	return coc.createWalletUseCase.Execute(ctx, email, &domain.Wallet{
		ID:         id,
		WalletName: wallet.Name,
		Balance:    0,
		OwnerID:    "",
	})
}
