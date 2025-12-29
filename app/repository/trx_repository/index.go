package trx_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type Trx_Repository struct {
	DB *gorm.DB
}

func NewTrxRepositoryRegistry() *Trx_Repository {
	return &Trx_Repository{
		DB: database.DB,
	}
}

func (repo *Trx_Repository) CreateTrx(trx *models.Trx) error {
	errCreate := repo.DB.Table("trxes").Create(trx).Error

	return errCreate
}

func (repo *Trx_Repository) CreateLog(log *models.LogProduk) error {
	errCreate := repo.DB.Table("log_produks").Create(log).Error

	return errCreate
}

func (repo *Trx_Repository) CreateDetailTrx(detail *[]models.DetailTrx) error {
	errCreate := repo.DB.Table("detail_trxes").Create(detail).Error

	return errCreate
}

func (repo *Trx_Repository) GetByIdTrx(IDTrx string) (*models.Trx, error) {

	var trx *models.Trx

	errGet := repo.DB.Table("trxes").Where("id = ?", IDTrx).First(&trx).Error

	return trx, errGet
}

func (repo *Trx_Repository) GetDetailByTrxId(IDTrx string) ([]models.DetailTrx, error) {
	var trx []models.DetailTrx

	errGet := repo.DB.Table("detail_trxes").Where("id_trx = ?", IDTrx).Find(&trx).Error

	return trx, errGet

}
