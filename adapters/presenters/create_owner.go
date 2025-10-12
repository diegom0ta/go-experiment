package presenters

import "experiment/adapters/presenters/input"

type CreateOwnerPresenter interface {
	Present(message string) *CreateOwnerResponse
}

type createOwnerPresenter struct{}

func NewCreateOwnerPresenter() *createOwnerPresenter {
	return &createOwnerPresenter{}
}

func (cop *createOwnerPresenter) Present(message string) *CreateOwnerResponse {
	return &CreateOwnerResponse{Message: message}
}

type CreateOwnerResponse struct {
	Message string `json:"message"`
}

type CreateOwnerRequest struct {
	Owner input.OwnerInput `json:"owner"`
}
