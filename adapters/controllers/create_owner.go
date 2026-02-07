package controllers

import (
	"context"
	"experiment/adapters/presenters/input"
	"experiment/core/domain"
	"experiment/usecases"

	"github.com/google/uuid"
)

type CreateOwnerController interface {
	HandleCreateOwner(ctx context.Context, owner *input.OwnerInput) error
}

type createOwnerController struct {
	createOwnerUseCase usecases.CreateOwnerUseCase
}

func NewCreateOwnerController(cou usecases.CreateOwnerUseCase) CreateOwnerController {
	return &createOwnerController{createOwnerUseCase: cou}
}

func (coc *createOwnerController) HandleCreateOwner(ctx context.Context, owner *input.OwnerInput) error {
	id := uuid.New().String()

	return coc.createOwnerUseCase.Execute(ctx, &domain.Owner{
		ID:        id,
		OwnerName: owner.Name,
		Email:     owner.Email,
		Document:  owner.Document,
		Wallets:   []domain.Wallet{},
	})
}
