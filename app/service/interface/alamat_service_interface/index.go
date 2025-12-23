package alamat_service_interface

import (
	"last-project/app/models"
	"last-project/app/request"
)

type Alamat_Service_Interface interface {
	Create(request *request.Alamat, IDUserParameter string) (*models.Alamat, error)
	GetByIDUserAndIDAlamat(IDUserParameter string, IDAlamat string) (*models.Alamat, error)
	GetByIDUser(IDUserParameter string) ([]models.Alamat, error)
	Update(IDUserParameter string, IDAlamat string, request *request.Alamat) (*models.Alamat, error)
	Delete(IDUserParameter string, IDAlamat string) error
}
