package services

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/dtos"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/helpers"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/models"
	r "git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/repositories"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type UserService interface {
	GetByEmail(email string) (*models.User, error)
	Login(user *dtos.Login) (*models.TokenResponse, error)
}

type userService struct {
	userRepository r.UserRepository
}

type USConfig struct {
	UserRepository r.UserRepository
}

func NewUserService(c *USConfig) UserService {
	return &userService{
		userRepository: c.UserRepository,
	}
}

func (u *userService) GetByEmail(email string) (*models.User, error) {
	return u.userRepository.FindUserByEmail(email)
}

func (u *userService) Login(user *dtos.Login) (*models.TokenResponse, error) {
	check, err := u.userRepository.FindUserByEmail(user.Email)
	if err != nil {
		return nil, helpers.FailedResponse(http.StatusBadRequest, "Invalid Credentials")
	}

	checkPassword := helpers.CheckPassword([]byte(check.Password), user.Password)
	if checkPassword != nil {
		return nil, helpers.FailedResponse(http.StatusBadRequest, "Invalid Credentials")
	}

	godotenv.Load("./.env")
	idExp, err := strconv.ParseInt(os.Getenv("TIME"), 10, 64)

	if err != nil {
		return nil, err
	}

	unixTime := time.Now().Unix()
	tokenExp := unixTime + idExp

	claims := &models.IdTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Entry Task MO FE",
			ExpiresAt: &jwt.NumericDate{Time: time.Unix(tokenExp, 0)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		User: &dtos.JwtToken{
			ID:       check.ID,
			FullName: check.Name,
			Role:     check.Role,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))
	if err != nil {
		return nil, err
	}
	return &models.TokenResponse{IDToken: tokenString, Role: check.Role, Name: check.Name, Email: check.Email, Photo: check.Photo}, nil
}
