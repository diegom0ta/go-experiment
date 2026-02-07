package presenters

import (
	"experiment/adapters/presenters/input"
	"experiment/adapters/presenters/output"
)

type CreateOwnerPresenter interface {
	Present(message string) *CreateOwnerResponse
}

type createOwnerPresenter struct{}

func NewCreateOwnerPresenter() CreateOwnerPresenter {
	return &createOwnerPresenter{}
}

func (cop *createOwnerPresenter) Present(message string) *CreateOwnerResponse {
	return &CreateOwnerResponse{Message: output.CreateOwnerOutput{Message: message}}
}

type CreateOwnerResponse struct {
	Message output.CreateOwnerOutput `json:"message"`
}

type CreateOwnerRequest struct {
	Owner input.OwnerInput `json:"owner"`
}
