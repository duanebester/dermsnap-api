package services

import (
	"dermsnap/models"
	"dermsnap/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DermsnapService interface {
	CreateDermsnap(userID uuid.UUID, opts models.CreateDermsnap) (*models.Dermsnap, error)
	CreateDermsnapImage(dermsnapID uuid.UUID, dermsnapImage *models.DermsnapImage) (*models.DermsnapImage, error)
	GetUserDermsnaps(userID uuid.UUID) ([]models.Dermsnap, error)
	GetDermsnapById(id uuid.UUID) (*models.Dermsnap, error)
	UpdateDermsnap(id uuid.UUID, dermsnap *models.Dermsnap) (*models.Dermsnap, error)
	DeleteDermsnap(dermsnap *models.Dermsnap) (*models.Dermsnap, error)
}

type DermsnapServiceImpl struct {
	dermsnapRepo repository.DermsnapRepository
}

func NewDermsnapService(db *gorm.DB) DermsnapService {
	return DermsnapServiceImpl{
		dermsnapRepo: repository.NewDermsnapRepository(db),
	}
}

func (d DermsnapServiceImpl) CreateDermsnap(userID uuid.UUID, opts models.CreateDermsnap) (*models.Dermsnap, error) {
	return d.dermsnapRepo.CreateDermsnap(userID, opts)
}

func (d DermsnapServiceImpl) GetDermsnapById(id uuid.UUID) (*models.Dermsnap, error) {
	return d.dermsnapRepo.GetDermsnapById(id)
}

func (d DermsnapServiceImpl) GetUserDermsnaps(userID uuid.UUID) ([]models.Dermsnap, error) {
	return d.dermsnapRepo.GetUserDermsnaps(userID)
}

func (d DermsnapServiceImpl) UpdateDermsnap(id uuid.UUID, opts *models.Dermsnap) (*models.Dermsnap, error) {
	return d.dermsnapRepo.UpdateDermsnap(id, opts)
}

func (d DermsnapServiceImpl) DeleteDermsnap(dermsnap *models.Dermsnap) (*models.Dermsnap, error) {
	return d.dermsnapRepo.DeleteDermsnap(dermsnap)
}

func (d DermsnapServiceImpl) CreateDermsnapImage(dermsnapID uuid.UUID, dermsnapImage *models.DermsnapImage) (*models.DermsnapImage, error) {
	return d.dermsnapRepo.CreateDermsnapImage(dermsnapID, dermsnapImage)
}
