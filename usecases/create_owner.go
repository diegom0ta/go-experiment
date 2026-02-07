package usecases

import (
	"context"
	"errors"
	"experiment/core/domain"
	"experiment/infra/logger"
	"experiment/ports"
)

var ErrOwnerAlreadyExists = errors.New("owner already exists")

type CreateOwnerUseCase interface {
	Execute(ctx context.Context, owner *domain.Owner) error
}

type createOwnerUseCase struct {
	ownerRepo  ports.OwnerRepository
	ownerCache ports.OwnerCache
}

func NewCreateOwnerUseCase(ownerRepo ports.OwnerRepository, ownerCache ports.OwnerCache) CreateOwnerUseCase {
	return &createOwnerUseCase{ownerRepo: ownerRepo, ownerCache: ownerCache}
}

func (couc *createOwnerUseCase) Execute(ctx context.Context, owner *domain.Owner) error {
	if existing, err := couc.ownerRepo.GetOwnerByEmail(owner.Email); err != nil {
		logger.Error("Error checking if owner exists: ", err)
		return err
	} else if existing != nil {
		logger.Warn("Owner already exists with email: ", owner.Email)
		return ErrOwnerAlreadyExists
	}

	err := couc.ownerCache.CacheOwner(ctx, owner)
	if err != nil {
		logger.Error("Error caching owner: ", err)
		return err
	}
	return nil
}
