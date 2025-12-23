package alamat_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type Alamat_Repository struct {
	DB *gorm.DB
}

func NewAlamatRepositoryRegistry() *Alamat_Repository {
	return &Alamat_Repository{
		DB: database.DB,
	}
}

func (repo *Alamat_Repository) Create(alamat *models.Alamat) error {
	errCreate := repo.DB.Table("alamats").Create(alamat).Error

	return errCreate
}

func (repo *Alamat_Repository) GetIDUserAndIDAlamat(IDUser string, IDAlamat string) (*models.Alamat, error) {
	var alamat *models.Alamat

	errGet := repo.DB.Table("alamats").Where("id_user = ? and id = ?", IDUser, IDAlamat).Find(&alamat).Error

	return alamat, errGet
}

func (repo *Alamat_Repository) GetIDUser(IDUser string) ([]models.Alamat, error) {
	var alamat []models.Alamat

	errGet := repo.DB.Table("alamats").Where("id_user = ?", IDUser).Find(&alamat).Error

	return alamat, errGet
}

func (repo *Alamat_Repository) Update(IDAlamat string, IDuser string, alamat *models.Alamat) error {
	errUpdate := repo.DB.Table("alamats").Where("id_user = ? and id = ?", IDuser, IDAlamat).Updates(alamat).Error

	return errUpdate
}

func (repo *Alamat_Repository) Delete(IDAlamat string, IDUser string) error {

	var alamat *models.Alamat

	errDelete := repo.DB.Table("alamats").Unscoped().Where("id = ?", IDAlamat).Delete(&alamat).Error

	return errDelete
}
