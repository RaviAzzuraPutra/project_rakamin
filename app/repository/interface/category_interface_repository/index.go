package category_interface_repository

import "last-project/app/models"

type Category_Interface_Repository interface {
	Create(category *models.Category) error
	Get() ([]models.Category, error)
	GetById(IDCategory string) (*models.Category, error)
	Update(IDCategory string, category *models.Category) error
	Delete(IDCategory string) error
}
