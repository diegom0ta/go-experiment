package usecases

import (
	"context"
	"experiment/core/domain"
	"experiment/ports"
)

type IGetWalletsByOwnerUseCase interface {
	Execute(ctx context.Context, email string) ([]*domain.Wallet, error)
}

type getWalletsByOwnerUseCase struct {
	walletRepo ports.WalletRepository
	ownerRepo  ports.OwnerRepository
	ownerCache ports.OwnerCache
}

func NewGetWalletsByOwnerUseCase(walletRepo ports.WalletRepository, ownerRepo ports.OwnerRepository, ownerCache ports.OwnerCache) IGetWalletsByOwnerUseCase {
	return &getWalletsByOwnerUseCase{walletRepo: walletRepo, ownerRepo: ownerRepo, ownerCache: ownerCache}
}

func (gwbouc *getWalletsByOwnerUseCase) Execute(ctx context.Context, email string) ([]*domain.Wallet, error) {
	owner, err := gwbouc.ownerCache.GetOwner(ctx, email)
	if err != nil {
		owner, err = gwbouc.ownerRepo.GetOwnerByEmail(email)
		if err != nil {
			return nil, err
		}
	}
	wallets, err := gwbouc.walletRepo.FindOwnerWallets(owner.ID)
	if err != nil {
		return nil, err
	}
	return wallets, nil
}
