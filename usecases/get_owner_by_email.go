package usecases

import (
	"context"
	"experiment/core/domain"
	"experiment/ports"
)

type GetOwnerByEmailUseCase interface {
	Execute(ctx context.Context, email string) (*domain.Owner, error)
}

type getOwnerByEmailUseCase struct {
	ownerRepo  ports.OwnerRepository
	ownerCache ports.OwnerCache
}

func NewGetOwnerByEmailUseCase(ownerRepo ports.OwnerRepository, ownerCache ports.OwnerCache) *getOwnerByEmailUseCase {
	return &getOwnerByEmailUseCase{ownerRepo: ownerRepo, ownerCache: ownerCache}
}

func (couc *getOwnerByEmailUseCase) Execute(ctx context.Context, email string) (*domain.Owner, error) {
	if owner, err := couc.ownerCache.GetOwner(ctx, email); err == nil {
		return owner, nil
	}
	return couc.ownerRepo.GetOwnerByEmail(email)
}
