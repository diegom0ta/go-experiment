package controllers

import (
	"experiment/usecases"
)

type DepositController interface {
	HandleDeposit(amount float64, currency, walletName, email string) error
}

type depositController struct {
	depositUseCase usecases.IDepositUseCase
}

func NewDepositController(du usecases.IDepositUseCase) DepositController {
	return &depositController{depositUseCase: du}
}

func (dc *depositController) HandleDeposit(amount float64, currency, walletName, email string) error {
	return dc.depositUseCase.Execute(amount, currency, walletName, email)
}
