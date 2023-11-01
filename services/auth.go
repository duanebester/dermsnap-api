package services

import (
	"dermsnap/repository"
	"dermsnap/utils"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"gorm.io/gorm"
)

type AuthService interface {
	LoginUser(email string, password string) (string, error)
}

type AuthServiceImpl struct {
	userRepo repository.UserRepository
}

func NewAuthService(db *gorm.DB) AuthService {
	return AuthServiceImpl{
		userRepo: repository.NewUserRepository(db),
	}
}

func (a AuthServiceImpl) LoginUser(email string, password string) (string, error) {
	user, err := a.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid password")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["iat"] = time.Now().Unix()
	claims["iss"] = "dermsnap"
	claims["sub"] = user.ID

	jwtSecret := os.Getenv("JWT_SECRET")
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}
