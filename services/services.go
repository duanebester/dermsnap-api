package services

import (
	"os"

	"gorm.io/gorm"
)

type Services struct {
	AuthService     AuthService
	UserService     UserService
	ImageService    ImageService
	DermsnapService DermsnapService
	ProviderService ProviderService
}

func NewServices(db *gorm.DB) *Services {
	var APP_ENV = os.Getenv("APP_ENV")
	var defaultImageBucketName = "dermsnap-dev"
	if APP_ENV == "production" {
		defaultImageBucketName = "dermsnap"
	}
	return &Services{
		AuthService:     NewAuthService(db),
		UserService:     NewUserService(db),
		ImageService:    NewImageService(defaultImageBucketName),
		DermsnapService: NewDermsnapService(db),
		ProviderService: NewProviderService(),
	}
}
