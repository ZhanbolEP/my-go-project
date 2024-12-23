package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/ZhanbolEP/my-go-project/models"
	"github.com/ZhanbolEP/my-go-project/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"golang.org/x/crypto/bcrypt"
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


	accessToken, err := s.generateToken(user, time.Hour*24)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.generateToken(user, user.Hour*24)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}


func (s *userService) generateToken(user *models.User, expiry time.Duration) (string, error) {
	now := time.Now()

	claims := JWTCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:	fmt.Sprintf("%d", user.ID),
			ExpiresAt:	jwt.NewNumericDate(now.Add(expiry)),
			IssuedAt:	jwt.NewNumericDate(now),
			ID: 		uuid.New().String(),
		},
		IsAdmin:	user.IsAdmin,
		Name:		user.Name,
		Email: 		user.Email,
		UserId: 	user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func (s *userService) RefreshToken(refreshToken string) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "" , errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", error.New("invalid token claims")
	}

	userId := uint(claims["sub"].(float64))
	user, err := s.userRepo.GetUserById(userId)
	if err != nil {
		return "", error.New("user not found")
	}

	return s.generateToken(user, time.Minute*15)
}

func (s *userService) VerifyToken(tokenString string) (*JWTCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error){
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTCustomClaims)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	if !token.Valid {
		return nil, errors.New("unauthorized")
	}
	return claims, nil
}