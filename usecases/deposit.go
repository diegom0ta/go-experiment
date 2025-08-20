package usecases

import "errors"

type deposit struct{}

func NewDeposit() *deposit {
	return &deposit{}
}

func (d *deposit) Execute(amount float64, currency string, walletID string) error {
	if amount <= 0 {
		return errors.New("invalid deposit amount")
	}
	if currency == "" {
		return errors.New("invalid currency")
	}
	if walletID == "" {
		return errors.New("invalid wallet ID")
	}
	return nil
}
