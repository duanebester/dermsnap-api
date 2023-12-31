package repository

import (
	"dermsnap/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id string) (*models.User, error)
	GetUserByIdentifier(identifier string, idType models.IdentifierType) (*models.User, error)
	CreateUser(identifier string, role models.Role, idType models.IdentifierType) (*models.User, error)
	CreateDoctorInfo(userID uuid.UUID, opts models.CreateDoctorInfo) (*models.DoctorInfo, error)
	GetUserInfo(userID uuid.UUID) (*models.UserInfo, error)
	GetDoctorInfo(userID uuid.UUID) (*models.DoctorInfo, error)
	CreateUserInfo(userID uuid.UUID, opts models.CreateUserInfo) (*models.UserInfo, error)
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

func (u UserRepositoryImpl) GetUserByIdentifier(identifier string, idType models.IdentifierType) (*models.User, error) {
	var user models.User
	err := u.db.Where("identifier = ? AND identifier_type = ?", identifier, idType).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// create user
func (u UserRepositoryImpl) CreateUser(identifier string, role models.Role, idType models.IdentifierType) (*models.User, error) {
	user := models.NewUser(identifier, role, idType)
	err := u.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// create doctor info
func (u UserRepositoryImpl) CreateDoctorInfo(userID uuid.UUID, opts models.CreateDoctorInfo) (*models.DoctorInfo, error) {
	doctorInfo := models.NewDoctorInfo(userID, opts.Specialty, opts.Credentials)
	err := u.db.Create(doctorInfo).Error
	if err != nil {
		return nil, err
	}
	return &doctorInfo, nil
}

func (u UserRepositoryImpl) GetUserInfo(userID uuid.UUID) (*models.UserInfo, error) {
	var userInfo models.UserInfo
	err := u.db.Where("user_id = ?", userID).First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func (u UserRepositoryImpl) GetDoctorInfo(userID uuid.UUID) (*models.DoctorInfo, error) {
	var doctorInfo models.DoctorInfo
	err := u.db.Where("user_id = ?", userID).First(&doctorInfo).Error
	if err != nil {
		return nil, err
	}
	return &doctorInfo, nil
}

func (u UserRepositoryImpl) CreateUserInfo(userID uuid.UUID, opts models.CreateUserInfo) (*models.UserInfo, error) {
	userInfo := models.NewUserInfo(userID, opts.Age, opts.Height, opts.Weight, opts.Gender)
	err := u.db.Create(userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}
