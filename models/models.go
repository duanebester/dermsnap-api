package models

import (
	"github.com/google/uuid"
)

type IdentifierType string

const (
	Doximity IdentifierType = "doximity"
	Apple    IdentifierType = "apple"
	Google   IdentifierType = "google"
)

type Role string

const (
	Admin  Role = "admin"
	Client Role = "client"
	Doctor Role = "doctor"
)

type User struct {
	ID             uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;"`
	Role           Role           `json:"role"`
	Identifier     string         `json:"identifier" gorm:"uniqueIndex:idx_identifier_type;not null;"`
	IdentifierType IdentifierType `json:"identifier_type" gorm:"uniqueIndex:idx_identifier_type;not null;"`
}

func NewUser(identifier string, role Role, idType IdentifierType) User {
	return User{
		ID:             uuid.New(),
		Role:           role,
		Identifier:     identifier,
		IdentifierType: idType,
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

func NewDoctorInfo(userID uuid.UUID, specialty, credentials string) DoctorInfo {
	return DoctorInfo{
		ID:          uuid.New(),
		UserID:      userID,
		Specialty:   specialty,
		Credentials: credentials,
	}
}
