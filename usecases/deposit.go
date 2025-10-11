package usecases

import "errors"

type deposit struct{}

func NewDeposit() *deposit {
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
