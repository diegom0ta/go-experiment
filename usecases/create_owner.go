package usecases

import (
	"experiment/core/domain"
	"experiment/ports"
)

type CreateOwnerUseCase interface {
	Execute(owner *domain.Owner) error
}

type createOwnerUseCase struct {
	ownerRepo ports.OwnerRepository
}

func NewCreateOwnerUseCase(ownerRepo ports.OwnerRepository) *createOwnerUseCase {
	return &createOwnerUseCase{ownerRepo: ownerRepo}
}

func (couc *createOwnerUseCase) Execute(owner *domain.Owner) error {
	return couc.ownerRepo.CreateOwner(owner)
}
