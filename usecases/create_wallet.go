package usecases

import (
	"context"
	"errors"
	"experiment/core/domain"
	"experiment/infra/logger"
	"experiment/ports"
)

var ErrWalletAlreadyExists = errors.New("wallet already exists")

type ICreateWalletUseCase interface {
	Execute(ctx context.Context, email string, wallet *domain.Wallet) error
}

type createWalletUseCase struct {
	walletRepo  ports.WalletRepository
	walletCache ports.WalletCache
	getOwner    getOwnerByEmailUseCase
}

func NewCreateWalletUseCase(walletRepo ports.WalletRepository, walletCache ports.WalletCache, getOwner getOwnerByEmailUseCase) ICreateWalletUseCase {
	return &createWalletUseCase{walletRepo: walletRepo, walletCache: walletCache, getOwner: getOwner}
}

func (cwuc *createWalletUseCase) Execute(ctx context.Context, email string, wallet *domain.Wallet) error {
	owner, err := cwuc.getOwner.Execute(ctx, email)
	if err != nil {
		logger.Error("Error checking if owner exists: ", err)
		return err
	}

	wallets, err := cwuc.walletRepo.FindOwnerWallets(owner.ID)
	if err != nil {
		logger.Error("Error retrieving owner's wallets: ", err)
		return err
	}

	for _, w := range wallets {
		if w.WalletName == wallet.WalletName {
			logger.Warn("Wallet already exists with name: ", wallet.WalletName)
			return ErrWalletAlreadyExists
		}
	}

	err = cwuc.walletRepo.CreateWallet(wallet)
	if err != nil {
		logger.Error("Error creating wallet: ", err)
		return err
	}

	return nil
}
