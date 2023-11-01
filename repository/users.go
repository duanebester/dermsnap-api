package repository

import (
	"dermsnap/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	RegisterUser(email string, password string) (*models.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepositoryImpl{
		db: db,
	}
}

func (u UserRepositoryImpl) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := u.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// create user
func (u UserRepositoryImpl) RegisterUser(email string, password string) (*models.User, error) {
	user, err := models.NewUser(models.User{
		Email:    email,
		Password: password,
		Role:     models.Client,
	})

	if err != nil {
		return nil, err
	}

	err = u.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
