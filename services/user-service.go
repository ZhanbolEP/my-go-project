package services

import (
	"github.com/ZhanbolEP/my-go-project/models"
	"github.com/ZhanbolEP/my-go-project/repositories"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService interface {
	RegisterUser(username, email, password string) error
	LoginUser(email, password string) (string, string, error)
	RefreshToken(refreshToken string) (string, error)
	VerifyToken(token string) (*JWTCustomClaims, error)
}

var jwtSecret = [] byte("AB12")

type JWTCustomClaims struct {
	jwt.RegisteredClaims
	IsAdmin		bool	`json:"is_admin"`
	Name		string	`json: "name"`
	Email		string	`json: "email"`
	UserId		uint	`json: "user_id"`
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{userRepo: repo}
}

//Password Hashing
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte{(password)})
	return err == nil
}

//Reg User
func (s *userService) RegisterUser(name, email, password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:		name,
		Email:		email,
		Password:	hashedPassword,
		IsAdmin:	false,
	}

	return s.userRepo.CreateUser(user)
}

//Login User
func (s *userService) LoginUser(email, password string) (string, string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", "", error.New("user not found")
	}

	if !checkPasswordHash(password, user.Password) {
		return "", "", error.New("incorrect password")
	}


	accessTokens
}