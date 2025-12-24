package user_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type User_Repository struct {
	DB *gorm.DB
}

func NewUserRepositoryRegistry() *User_Repository {
	return &User_Repository{
		DB: database.DB,
	}
}

func (repo *User_Repository) Get(IDUser string) (*models.User, error) {
	var user *models.User

	errGet := repo.DB.Table("users").Where("id = ?", IDUser).First(&user).Error

	return user, errGet
}

func (repo *User_Repository) Update(IDUser string, user *models.User) error {
	errUpdate := repo.DB.Table("users").Where("id = ?", IDUser).Updates(&user).Error

	return errUpdate
}
