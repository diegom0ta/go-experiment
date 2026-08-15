package usecases

import "errors"

type deposit struct{}

type IDepositUseCase interface {
	Execute(amount float64, currency string, walletName string) error
}

func NewDepositUseCase() IDepositUseCase {
	return &deposit{}
}

func (d *deposit) Execute(amount float64, currency string, walletName string) error {
	if amount <= 0 {
		return errors.New("invalid deposit amount")
	}
	if currency == "" {
		return errors.New("invalid currency")
	}
	if walletName == "" {
		return errors.New("invalid wallet name")
	}
	return nil
}
