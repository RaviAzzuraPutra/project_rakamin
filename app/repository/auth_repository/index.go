package auth_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type Auth_Repository struct {
	DB *gorm.DB
}

func NewAuthRpositoryRegistry() *Auth_Repository {
	return &Auth_Repository{
		DB: database.DB,
	}
}

func (repo *Auth_Repository) Register(register *models.User) error {
	errRegister := repo.DB.Table("users").Create(register).Error

	return errRegister
}

func (repo *Auth_Repository) Login(email string) (*models.User, error) {
	var user models.User

	errLogin := repo.DB.Table("users").Where("email = ?", email).First(&user).Error

	return &user, errLogin
}
