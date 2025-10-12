package repository

import (
	"errors"

	"experiment/core/domain"
	"experiment/infra/database"

	"gorm.io/gorm"
)

type WalletRepository struct{}

func NewWalletRepository() *WalletRepository {
	return &WalletRepository{}
}

func (r *WalletRepository) CreateWallet(wallet *domain.Wallet) error {
	return database.DB.Create(wallet).Error
}

func (r *WalletRepository) FindOwnerWallets(ownerID string) ([]domain.Wallet, error) {
	var wallets []domain.Wallet
	result := database.DB.Where("owner_id = ?", ownerID).Find(&wallets)
	return wallets, result.Error
}

// FindWalletByName finds a wallet by its name
func (r *WalletRepository) FindWalletByName(name string) (*domain.Wallet, error) {
	var wallet domain.Wallet
	result := database.DB.Where("name = ?", name).First(&wallet)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("wallet not found")
	}
	return &wallet, result.Error
}

func (r *WalletRepository) GetWalletByID(walletID string) (*domain.Wallet, error) {
	var wallet domain.Wallet
	result := database.DB.First(&wallet, "id = ?", walletID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &wallet, nil
}

func (r *WalletRepository) GetAllWallets() ([]domain.Wallet, error) {
	var wallets []domain.Wallet
	result := database.DB.Find(&wallets)
	return wallets, result.Error
}

func (r *WalletRepository) DeleteWallet(walletID string) error {
	return database.DB.Delete(&domain.Wallet{}, "id = ?", walletID).Error
}

func (r *WalletRepository) UpdateWalletByName(wallet *domain.Wallet) error {
	walletInDB, err := r.FindWalletByName(wallet.WalletName)
	if err != nil {
		return err
	}
	if walletInDB == nil {
		return errors.New("wallet not found")
	}
	wallet.ID = walletInDB.ID
	return database.DB.Save(wallet).Error
}

// Deposit adds the amount to the wallet's balance and saves it
func (r *WalletRepository) Deposit(walletName string, amount int) error {
	wallet, err := r.FindWalletByName(walletName)
	if err != nil {
		return err
	}
	if amount <= 0 {
		return errors.New("invalid deposit amount")
	}
	wallet.Balance += amount
	return database.DB.Save(wallet).Error
}

func (r *WalletRepository) Withdraw(walletName string, amount int) error {
	wallet, err := r.FindWalletByName(walletName)
	if err != nil {
		return err
	}
	if amount <= 0 {
		return errors.New("invalid withdraw amount")
	}
	if wallet.Balance < amount {
		return errors.New("insufficient funds")
	}
	wallet.Balance -= amount
	return database.DB.Save(wallet).Error
}
