package services

import (
	"gorm.io/gorm"
)

type Services struct {
	AuthService     AuthService
	UserService     UserService
	DermsnapService DermsnapService
	ProviderService ProviderService
}

func NewServices(db *gorm.DB) *Services {
	return &Services{
		AuthService:     NewAuthService(db),
		UserService:     NewUserService(db),
		DermsnapService: NewDermsnapService(db),
		ProviderService: NewProviderService(),
	}
}
