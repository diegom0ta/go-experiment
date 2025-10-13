package presenters

import (
	"experiment/adapters/presenters/output"
)

type GetOwnerPresenter interface {
	Present(owner *output.GetOwnerOutput) *GetOwnerResponse
}

type getOwnerPresenter struct{}

func NewGetOwnerPresenter() *getOwnerPresenter {
	return &getOwnerPresenter{}
}

func (gop *getOwnerPresenter) Present(owner *output.GetOwnerOutput) *GetOwnerResponse {
	return &GetOwnerResponse{Owner: *owner}
}

type GetOwnerResponse struct {
	Owner output.GetOwnerOutput `json:"owner"`
}
