package services

import (
	"dermsnap/models"
	"dermsnap/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService interface {
	GetUserByID(id string) (*models.User, error)
	GetUserByIdentifier(identifier string, idType models.IdentifierType) (*models.User, error)
	CreateUser(identifier string, role models.Role, idType models.IdentifierType) (*models.User, error)
	GetUserInfo(userID uuid.UUID) (*models.UserInfo, error)
	CreateUserInfo(userID uuid.UUID, opts models.CreateUserInfo) (*models.UserInfo, error)
	CreateDoctorInfo(userID uuid.UUID, specialty string, credentials string) (*models.DoctorInfo, error)
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

func (u UserServiceImpl) CreateUser(identifier string, role models.Role, idType models.IdentifierType) (*models.User, error) {
	return u.userRepo.CreateUser(identifier, role, idType)
}

func (u UserServiceImpl) GetUserByIdentifier(identifier string, idType models.IdentifierType) (*models.User, error) {
	return u.userRepo.GetUserByIdentifier(identifier, idType)
}

func (u UserServiceImpl) CreateDoctorInfo(userID uuid.UUID, specialty string, credentials string) (*models.DoctorInfo, error) {
	return u.userRepo.CreateDoctorInfo(userID, specialty, credentials)
}

func (u UserServiceImpl) GetUserInfo(userID uuid.UUID) (*models.UserInfo, error) {
	return u.userRepo.GetUserInfo(userID)
}

func (u UserServiceImpl) CreateUserInfo(userID uuid.UUID, opts models.CreateUserInfo) (*models.UserInfo, error) {
	return u.userRepo.CreateUserInfo(userID, opts)
}
