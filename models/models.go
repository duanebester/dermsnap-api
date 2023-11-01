package models

import (
	"dermsnap/utils"

	"github.com/google/uuid"
)

type Role string

const (
	Admin  Role = "admin"
	Client Role = "client"
	Doctor Role = "doctor"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"-"`
	Role     Role      `json:"role"`
}

func NewUser(opts User) (*User, error) {
	hashedPassword, err := utils.HashPassword(opts.Password)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       uuid.New(),
		Email:    opts.Email,
		Password: hashedPassword,
		Role:     opts.Role,
	}, nil
}
