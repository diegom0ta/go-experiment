package repository

import (
	"database/sql"

	"experiment/core/domain"
	"experiment/infra/database"
	"experiment/infra/logger"

	"github.com/sirupsen/logrus"
)

type OwnerRepository struct{}

func NewOwnerRepository() *OwnerRepository {
	return &OwnerRepository{}
}

func (r *OwnerRepository) CreateOwner(owner *domain.Owner) error {
	logger.WithFields(logrus.Fields{
		"owner_id": owner.ID,
		"email":    owner.Email,
	}).Info("Creating new owner")

	err := database.DB.Create(owner).Error
	if err != nil {
		logger.WithFields(logrus.Fields{
			"owner_id": owner.ID,
			"error":    err.Error(),
		}).Error("Failed to create owner")
		return err
	}

	logger.WithField("owner_id", owner.ID).Info("Owner created successfully")
	return nil
}

func (r *OwnerRepository) GetOwnerByID(ownerID string) (*domain.Owner, error) {
	logger.WithField("owner_id", ownerID).Debug("Fetching owner by ID")

	var owner domain.Owner
	result := database.DB.First(&owner, "id = ?", ownerID)
	if result.Error != nil {
		if result.Error == sql.ErrNoRows {
			logger.WithField("owner_id", ownerID).Warn("Owner not found")
			return nil, nil
		}
		logger.WithFields(logrus.Fields{
			"owner_id": ownerID,
			"error":    result.Error.Error(),
		}).Error("Failed to fetch owner")
		return nil, result.Error
	}

	logger.WithField("owner_id", ownerID).Debug("Owner fetched successfully")
	return &owner, nil
}

func (r *OwnerRepository) GetAllOwners() ([]domain.Owner, error) {
	var owners []domain.Owner
	result := database.DB.Find(&owners)
	return owners, result.Error
}

func (r *OwnerRepository) DeleteOwner(ownerID string) error {
	logger.WithField("owner_id", ownerID).Info("Deleting owner")

	err := database.DB.Delete(&domain.Owner{}, "id = ?", ownerID).Error
	if err != nil {
		logger.WithFields(logrus.Fields{
			"owner_id": ownerID,
			"error":    err.Error(),
		}).Error("Failed to delete owner")
		return err
	}

	logger.WithField("owner_id", ownerID).Info("Owner deleted successfully")
	return nil
}

func (r *OwnerRepository) UpdateOwner(owner *domain.Owner) error {
	return database.DB.Save(owner).Error
}
