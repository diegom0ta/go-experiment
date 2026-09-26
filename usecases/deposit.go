package usecases

import (
	"errors"
	"experiment/ports"
)

type deposit struct {
	ownerRepo ports.OwnerRepository
	walletRepo ports.WalletRepository
}

type IDepositUseCase interface {
	Execute(amount float64, currency, walletName, email string) error
}

func NewDepositUseCase(ownerRepo ports.OwnerRepository, walletRepo ports.WalletRepository) IDepositUseCase {
	return &deposit{
		ownerRepo: ownerRepo,
		walletRepo: walletRepo,
	}
}

func (d *deposit) Execute(amount float64, currency, walletName, email string) error {
	if amount <= 0 {



		return errors.New("invalid deposit amount")
	}
	if currency == "" {
		return errors.New("invalid currency")
	}
	if walletName == "" {
		return errors.New("invalid wallet name")
	}

	owner, err := d.ownerRepo.GetOwnerByEmail(email)
	if err != nil {
		return err
	}
	if owner == nil {
		return errors.New("owner not found")
	}

	wallet, err := d.walletRepo.FindWalletByName(walletName)
	if err != nil {
		return err
	}
	if wallet == nil {
		return errors.New("wallet not found")
	}
	if wallet.OwnerID != owner.ID {
		return errors.New("wallet does not belong to the owner")
	}

	a := int64(amount * 100) // Convert to cents
	err = d.walletRepo.Deposit(walletName, a)
	if err != nil {
		return err
	}

	return nil
}
