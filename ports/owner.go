package ports

import (
	"context"
	"experiment/core/domain"
)

type OwnerRepository interface {
	CreateOwner(owner *domain.Owner) error
	GetOwnerByEmail(email string) (*domain.Owner, error)
	GetAllOwners() ([]domain.Owner, error)
	DeleteOwner(ownerID string) error
	UpdateOwner(owner *domain.Owner) error
}

type OwnerCache interface {
	CacheOwner(ctx context.Context, owner *domain.Owner) error
	GetOwner(ctx context.Context, email string) (*domain.Owner, error)
	DeleteOwner(ctx context.Context, email string) error
}
