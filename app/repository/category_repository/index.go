package category_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type Category_Repository struct {
	DB *gorm.DB
}

func NewCategoryRepositoryRegistry() *Category_Repository {
	return &Category_Repository{
		DB: database.DB,
	}
}

func (repo *Category_Repository) Create(category *models.Category) error {
	errCreate := repo.DB.Table("categories").Create(category).Error

	return errCreate
}

func (repo *Category_Repository) Get() ([]models.Category, error) {

	var category []models.Category

	errGet := repo.DB.Table("categories").Find(&category).Error

	return category, errGet
}

func (repo *Category_Repository) GetById(IDCategory string) (*models.Category, error) {

	var category *models.Category

	errGet := repo.DB.Table("categories").Where("id = ?", IDCategory).First(&category).Error

	return category, errGet
}

func (repo *Category_Repository) Update(IDCategory string, category *models.Category) error {
	errUpdate := repo.DB.Table("categories").Where("id = ?", IDCategory).Updates(category).Error

	return errUpdate
}

func (repo *Category_Repository) Delete(IDCategory string) error {
	var category *models.Category

	errDelete := repo.DB.Table("categories").Unscoped().Where("id = ?", IDCategory).Delete(&category).Error

	return errDelete
}
