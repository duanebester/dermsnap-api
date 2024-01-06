package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
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
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

func NewUser(identifier string, role Role, idType IdentifierType) User {
	return User{
		ID:             uuid.New(),
		Role:           role,
		Identifier:     identifier,
		IdentifierType: idType,
	}
}

type CreateUserInfo struct {
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

type UserInfo struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;index:idx_user_id,unique"`
	Height int       `json:"height"`
	Weight int       `json:"weight"`
	Age    int       `json:"age"`
	Gender string    `json:"gender"`
}

func NewUserInfo(userID uuid.UUID, age, height, weight int, gender string) UserInfo {
	return UserInfo{
		ID:     uuid.New(),
		UserID: userID,
		Height: height,
		Weight: weight,
		Age:    age,
		Gender: gender,
	}
}

type CreateDoctorInfo struct {
	Specialty   string `json:"specialty"`
	Credentials string `json:"credentials"`
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

type BodyLocation string

const (
	Scalp    BodyLocation = "scalp"
	Face     BodyLocation = "face" // (not including eyes or mouth)
	Eyes     BodyLocation = "eyes"
	Mouth    BodyLocation = "mouth"
	Neck     BodyLocation = "neck"
	Chest    BodyLocation = "chest"
	Abdomen  BodyLocation = "abdomen"
	Back     BodyLocation = "back"
	Arms     BodyLocation = "arms"
	Hands    BodyLocation = "hands"
	Buttocks BodyLocation = "buttocks"
	Genitals BodyLocation = "genitals"
	Legs     BodyLocation = "legs"
	Feet     BodyLocation = "feet"
)

type CreateDermsnap struct {
	StartTime      time.Time      `json:"start_time"`
	Duration       int            `json:"duration"`
	Locations      []BodyLocation `json:"locations"`
	Changed        bool           `json:"changed"`
	NewMedications []string       `json:"new_medications"`
	Itchy          bool           `json:"itchy"`
	Painful        bool           `json:"painful"`
	MoreInfo       string         `json:"more_info"`
}

type UpdateDermsnap struct {
	StartTime      time.Time      `json:"start_time"`
	Duration       int            `json:"duration"`
	Locations      []BodyLocation `json:"locations"`
	Changed        bool           `json:"changed"`
	NewMedications []string       `json:"new_medications"`
	Itchy          bool           `json:"itchy"`
	Painful        bool           `json:"painful"`
	MoreInfo       string         `json:"more_info"`
}

type DermsnapImage struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	DermsnapID uuid.UUID `json:"dermsnap_id" gorm:"type:uuid"`
	ImagePath  string    `json:"image_path"`
	CreatedAt  time.Time `json:"created_at"`
}

type Dermsnap struct {
	ID             uuid.UUID       `json:"id" gorm:"type:uuid;primary_key;"`
	UserID         uuid.UUID       `json:"user_id" gorm:"type:uuid;"`
	Reviewed       bool            `json:"reviewed" gorm:"default:false"`
	ReviewedBy     uuid.UUID       `json:"reviewed_by,omitempty" gorm:"type:uuid;default:null;"`
	StartTime      time.Time       `json:"start_time"`
	Duration       int             `json:"duration"`
	Locations      pq.StringArray  `json:"locations" gorm:"type:varchar(64)[]"`
	Changed        bool            `json:"changed"`
	NewMedications pq.StringArray  `json:"new_medications" gorm:"type:varchar(255)[];"`
	Itchy          bool            `json:"itchy"`
	Painful        bool            `json:"painful"`
	MoreInfo       string          `json:"more_info"`
	Images         []DermsnapImage `json:"images"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

func NewDermsnap(userID uuid.UUID, opts CreateDermsnap) Dermsnap {
	locations := make([]string, len(opts.Locations))
	for i, location := range opts.Locations {
		locations[i] = string(location)
	}
	return Dermsnap{
		ID:             uuid.New(),
		UserID:         userID,
		StartTime:      opts.StartTime,
		Duration:       opts.Duration,
		Locations:      locations,
		Changed:        opts.Changed,
		NewMedications: opts.NewMedications,
		Itchy:          opts.Itchy,
		Painful:        opts.Painful,
		MoreInfo:       opts.MoreInfo,
	}
}
