package category_service_interface

import (
	"last-project/app/models"
	"last-project/app/request"
)

type Category_Service_Interface interface {
	Create(request *request.Category_Request) (*models.Category, error)
	Get() ([]models.Category, error)
	GetById(IDCategory string) (*models.Category, error)
	Update(IDCategory string, request *request.Category_Request) (*models.Category, error)
	Delete(IDCategory string) error
}
