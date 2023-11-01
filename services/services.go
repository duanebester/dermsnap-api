package services

import (
	"gorm.io/gorm"
)

type Services struct {
	UserService UserService
	AuthService AuthService
}

func NewServices(db *gorm.DB) *Services {
	return &Services{
		UserService: NewUserService(db),
		AuthService: NewAuthService(db),
	}
}
