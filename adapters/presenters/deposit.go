package presenters

import "experiment/adapters/presenters/output"

type DepositPresenter interface {
	Present(deposit *output.DepositOutput) *DepositResponse
}

type depositPresenter struct{}

func NewDepositPresenter() DepositPresenter {
	return &depositPresenter{}
}

func (dp *depositPresenter) Present(deposit *output.DepositOutput) *DepositResponse {
	return &DepositResponse{Deposit: *deposit}
}

type DepositResponse struct {
	Deposit output.DepositOutput `json:"deposit"`
}
