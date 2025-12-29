package foto_service_interface

import (
	"last-project/app/models"
	"last-project/app/request"
)

type Foto_Service_Interface interface {
	AddPhoto(request *request.Foto_Product_Request, IDUser string, IDProduct string) (*models.FotoProduk, error)
	Delete(IDFoto string) error
	GetById(IDFoto string) (*models.FotoProduk, error)
	GetByIdProduct(IDProduct string) ([]models.FotoProduk, error)
}
