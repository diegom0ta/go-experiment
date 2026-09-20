package controllers

import (
	"context"
	"experiment/core/domain"
	"experiment/usecases"
)

type GetWalletsByOwnerController interface {
	HandleGetWalletsByOwner(ctx context.Context, email string) ([]*domain.Wallet, error)
}

type getWalletsByOwnerController struct {
	getWalletsByOwnerUseCase usecases.IGetWalletsByOwnerUseCase
}

func NewGetWalletsByOwnerController(gwbouc usecases.IGetWalletsByOwnerUseCase) GetWalletsByOwnerController {
	return &getWalletsByOwnerController{getWalletsByOwnerUseCase: gwbouc}
}

func (gwbouc *getWalletsByOwnerController) HandleGetWalletsByOwner(ctx context.Context, email string) ([]*domain.Wallet, error) {
	return gwbouc.getWalletsByOwnerUseCase.Execute(ctx, email)
}
