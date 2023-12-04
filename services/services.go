package services

import (
	"gorm.io/gorm"
)

type Services struct {
	AuthService     AuthService
	UserService     UserService
	ProviderService ProviderService
}

func NewServices(db *gorm.DB) *Services {
	return &Services{
		AuthService:     NewAuthService(db),
		UserService:     NewUserService(db),
		ProviderService: NewProviderService(),
	}
}
