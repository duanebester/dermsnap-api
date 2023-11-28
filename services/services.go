package services

import (
	"gorm.io/gorm"
)

type Services struct {
	UserService UserService
}

func NewServices(db *gorm.DB) *Services {
	return &Services{
		UserService: NewUserService(db),
	}
}
