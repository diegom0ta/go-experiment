package repository

import (
	"errors"

	"experiment/infra/database"

	"gorm.io/gorm"
)

type Wallet struct {
	ID      string
	Name    string
	Balance int
}

type DepositRepository struct{}

func NewDepositRepository() *DepositRepository {
	return &DepositRepository{}
}

// FindWalletByName finds a wallet by its name
func (r *DepositRepository) FindWalletByName(name string) (*Wallet, error) {
	var wallet Wallet
	result := database.DB.Where("name = ?", name).First(&wallet)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("wallet not found")
	}
	return &wallet, result.Error
}

// Deposit adds the amount to the wallet's balance and saves it
func (r *DepositRepository) Deposit(walletName string, amount int) error {
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
