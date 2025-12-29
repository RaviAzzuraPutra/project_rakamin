package product_interface_repository

import "last-project/app/models"

type Product_Interface_Repository interface {
	Create(product *models.Produk) error
	GetAll() ([]models.Produk, error)
	GetByIdAndIdToko(IDProduct string, IDToko string) (*models.Produk, error)
	GetByCategory(search string) ([]models.Produk, error)
	Update(IDProduct string, IDToko string, Product *models.Produk) error
	Delete(IDProduct string, IDToko string) error
	GetById(IDProduct string) (*models.Produk, error)
	GetByIdToko(IDToko string) ([]models.Produk, error)
	ReduceStock(IDProduct string, IDToko string, qty int) error
}
