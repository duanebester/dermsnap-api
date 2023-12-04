package services

import (
	"dermsnap/models"
	"dermsnap/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService interface {
	GenerateToken(email string, role models.Role, idType models.IdentifierType) (string, error)
}

type AuthServiceImpl struct {
	userRepo repository.UserRepository
}

func NewAuthService(db *gorm.DB) AuthService {
	return AuthServiceImpl{
		userRepo: repository.NewUserRepository(db),
	}
}

func (a AuthServiceImpl) GenerateToken(identifier string, role models.Role, idType models.IdentifierType) (string, error) {
	user, err := a.userRepo.GetUserByIdentifier(identifier, idType)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["id_type"] = string(idType)
	claims["role"] = string(role)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	jwtSecret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(jwtSecret))
}
