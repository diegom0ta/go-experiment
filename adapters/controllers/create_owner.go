package controllers

import (
	"experiment/adapters/presenters/input"
	"experiment/core/domain"
	"experiment/usecases"

	"github.com/google/uuid"
)

type CreateOwnerController interface {
	HandleCreateOwner(owner *input.OwnerInput) error
}

type createOwnerController struct {
	createOwnerUseCase usecases.CreateOwnerUseCase
}

func NewCreateOwnerController(cou usecases.CreateOwnerUseCase) *createOwnerController {
	return &createOwnerController{createOwnerUseCase: cou}
}

func (coc *createOwnerController) HandleCreateOwner(owner *input.OwnerInput) error {
	uuid := uuid.New().String()

	return coc.createOwnerUseCase.Execute(&domain.Owner{
		ID:        uuid,
		OwnerName: owner.Name,
		Email:     owner.Email,
		Document:  owner.Document,
		Wallets:   []domain.Wallet{},
	})
}
