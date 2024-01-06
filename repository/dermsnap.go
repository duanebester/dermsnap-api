package repository

import (
	"dermsnap/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DermsnapRepository interface {
	CreateDermsnap(userID uuid.UUID, opts models.CreateDermsnap) (*models.Dermsnap, error)
	GetUserDermsnaps(userID uuid.UUID) ([]models.Dermsnap, error)
	GetDermsnapById(id uuid.UUID) (*models.Dermsnap, error)
	UpdateDermsnap(id uuid.UUID, opts models.UpdateDermsnap) (*models.Dermsnap, error)
	DeleteDermsnap(dermsnap *models.Dermsnap) (*models.Dermsnap, error)
}

type DermsnapRepositoryImpl struct {
	db *gorm.DB
}

func NewDermsnapRepository(db *gorm.DB) DermsnapRepository {
	return DermsnapRepositoryImpl{
		db: db,
	}
}

func (d DermsnapRepositoryImpl) CreateDermsnap(userID uuid.UUID, opts models.CreateDermsnap) (*models.Dermsnap, error) {
	dermsnap := models.NewDermsnap(userID, opts)
	err := d.db.Create(&dermsnap).Error
	if err != nil {
		return nil, err
	}
	return &dermsnap, nil
}

func (d DermsnapRepositoryImpl) GetDermsnapById(id uuid.UUID) (*models.Dermsnap, error) {
	var dermsnap models.Dermsnap
	err := d.db.Where("id = ?", id).First(&dermsnap).Error
	if err != nil {
		return nil, err
	}
	return &dermsnap, nil
}

func (d DermsnapRepositoryImpl) UpdateDermsnap(id uuid.UUID, opts models.UpdateDermsnap) (*models.Dermsnap, error) {
	err := d.db.Where("id = ?", id).Updates(&opts).Error
	if err != nil {
		return nil, err
	}
	return d.GetDermsnapById(id)
}

func (d DermsnapRepositoryImpl) GetUserDermsnaps(userID uuid.UUID) ([]models.Dermsnap, error) {
	var dermsnaps []models.Dermsnap
	err := d.db.Where("user_id = ?", userID).Find(&dermsnaps).Error
	if err != nil {
		return nil, err
	}
	return dermsnaps, nil
}

func (d DermsnapRepositoryImpl) DeleteDermsnap(dermsnap *models.Dermsnap) (*models.Dermsnap, error) {
	err := d.db.Delete(&dermsnap).Error
	if err != nil {
		return nil, err
	}
	return dermsnap, nil
}
