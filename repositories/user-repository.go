package repositories

import (
	"github.com/ZhanbolEP/my-go-project/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserById(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userRepostory struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepostory{db}
}

func (r *userRepostory) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepostory) GetUserById(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *userRepostory) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}