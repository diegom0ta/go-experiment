package ports

import (
	"context"
	"experiment/core/domain"
)

type WalletRepository interface {
	CreateWallet(wallet *domain.Wallet) error
	FindOwnerWallets(ownerID string) ([]domain.Wallet, error)
	FindWalletByName(name string) (*domain.Wallet, error)
	GetWalletByID(walletID string) (*domain.Wallet, error)
	GetAllWallets() ([]domain.Wallet, error)
	DeleteWallet(walletID string) error
	UpdateWalletByName(wallet *domain.Wallet) error
	Deposit(walletName string, amount int) error
	Withdraw(walletName string, amount int) error
}

type WalletCache interface {
	CacheWallet(ctx context.Context, wallet *domain.Wallet) error
	GetWallet(ctx context.Context, walletID string) (*domain.Wallet, error)
	DeleteWallet(ctx context.Context, walletID string) error
}
