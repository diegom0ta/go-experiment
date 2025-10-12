package repository

import (
	"database/sql"

	"experiment/core/domain"
	"experiment/infra/database"
)

type OwnerRepository struct{}

func NewOwnerRepository() *OwnerRepository {
	return &OwnerRepository{}
}

func (r *OwnerRepository) CreateOwner(owner *domain.Owner) error {
	return database.DB.Create(owner).Error
}

func (r *OwnerRepository) GetOwnerByID(ownerID string) (*domain.Owner, error) {
	var owner domain.Owner
	result := database.DB.First(&owner, "id = ?", ownerID)
	if result.Error != nil {
		if result.Error == sql.ErrNoRows {
			return nil, nil
		}
		return nil, result.Error
	}
	return &owner, nil
}

func (r *OwnerRepository) GetAllOwners() ([]domain.Owner, error) {
	var owners []domain.Owner
	result := database.DB.Find(&owners)
	return owners, result.Error
}

func (r *OwnerRepository) DeleteOwner(ownerID string) error {
	return database.DB.Delete(&domain.Owner{}, "id = ?", ownerID).Error
}

func (r *OwnerRepository) UpdateOwner(owner *domain.Owner) error {
	return database.DB.Save(owner).Error
}
