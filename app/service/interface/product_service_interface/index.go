package product_service_interface

import (
	"last-project/app/models"
	"last-project/app/request"
)

type Product_Service_Interface interface {
	Create(request *request.Product_Request, IDUser string) (*models.Produk, error)
	GetAll() ([]models.Produk, error)
	GetByIdAndIdToko(IDProduct string, IDUser string) (*models.Produk, error)
	GetByCategory(search string) ([]models.Produk, error)
	Update(IDProduct string, IDUser string, request *request.Product_Request) (*models.Produk, error)
	Delete(IDProduct string, IDUser string) error
	GetById(IDProduct string) (*models.Produk, error)
	GetByIdToko(IDUser string) ([]models.Produk, error)
}
