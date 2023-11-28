package models

import (
	"github.com/google/uuid"
)

type IdentifierType string

const (
	Email    IdentifierType = "email"
	Doximity IdentifierType = "doximity"
	Apple    IdentifierType = "apple"
	Google   IdentifierType = "google"
)

type Identifier struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;"`
	UserID     uuid.UUID      `json:"user_id" gorm:"type:uuid;"`
	Identifier string         `json:"identifier"`
	Type       IdentifierType `json:"type"`
}

type Role string

const (
	Admin  Role = "admin"
	Client Role = "client"
	Doctor Role = "doctor"
)

type User struct {
	ID       uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;"`
	Identity Identifier `json:"identity,inline"`
	Role     Role       `json:"role"`
}

func NewUser(email string, role Role) User {
	return User{
		ID: uuid.New(),
		Identity: Identifier{
			ID:         uuid.New(),
			Identifier: email,
			Type:       Email,
		},
		Role: role,
	}
}

type UserInfo struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid"`
	Height int       `json:"height"`
	Weight int       `json:"weight"`
	Age    int       `json:"age"`
	Gender string    `json:"gender"`
}

type DoctorInfo struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid"`
	Specialty   string    `json:"specialty"`
	Credentials string    `json:"credentials"`
}
