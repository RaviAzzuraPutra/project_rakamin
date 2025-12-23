package alamat_interface_repository

import "last-project/app/models"

type Alamat_Interface_Repository interface {
	Create(alamat *models.Alamat) error
	GetIDUserAndIDAlamat(IDUser string, IDAlamat string) (*models.Alamat, error)
	Update(IDAlamat string, IDuser string, alamat *models.Alamat) error
	Delete(IDAlamat string, IDUser string) error
	GetIDUser(IDUser string) ([]models.Alamat, error)
}
