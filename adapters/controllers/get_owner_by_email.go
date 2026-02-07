package controllers

import (
	"context"
	"experiment/core/domain"
	"experiment/usecases"
)

type GetOwnerByEmailController interface {
	HandleGetOwnerByEmail(ctx context.Context, email string) (*domain.Owner, error)
}

type getOwnerByEmailController struct {
	getOwnerByEmailUseCase usecases.GetOwnerByEmailUseCase
}

func NewGetOwnerByEmailController(gou usecases.GetOwnerByEmailUseCase) GetOwnerByEmailController {
	return &getOwnerByEmailController{getOwnerByEmailUseCase: gou}
}

func (coc *getOwnerByEmailController) HandleGetOwnerByEmail(ctx context.Context, email string) (*domain.Owner, error) {
	return coc.getOwnerByEmailUseCase.Execute(ctx, email)
}
