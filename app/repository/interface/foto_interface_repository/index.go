package foto_interface_repository

import "last-project/app/models"

type Foto_Interface_Repository interface {
	AddPhoto(foto *models.FotoProduk) error
	DeleteFoto(IDFoto string) error
	GetById(IDFoto string) (*models.FotoProduk, error)
	GetByIdProduct(IDProduct string) ([]models.FotoProduk, error)
}
