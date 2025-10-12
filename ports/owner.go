package ports

import "experiment/core/domain"

type OwnerRepository interface {
	CreateOwner(owner *domain.Owner) error
	GetOwnerByID(ownerID string) (*domain.Owner, error)
	GetAllOwners() ([]domain.Owner, error)
	DeleteOwner(ownerID string) error
	UpdateOwner(owner *domain.Owner) error
}
