package services

import (
	"dermsnap/models"
	"dermsnap/repository"

	"gorm.io/gorm"
)

type UserService interface {
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	RegisterUser(email string, password string) (*models.User, error)
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(db *gorm.DB) UserService {
	return UserServiceImpl{
		userRepo: repository.NewUserRepository(db),
	}
}

func (u UserServiceImpl) GetUserByID(id string) (*models.User, error) {
	return u.userRepo.GetUserByID(id)
}

func (u UserServiceImpl) GetUserByEmail(email string) (*models.User, error) {
	return u.userRepo.GetUserByEmail(email)
}

func (u UserServiceImpl) RegisterUser(email string, password string) (*models.User, error) {
	return u.userRepo.RegisterUser(email, password)
}
